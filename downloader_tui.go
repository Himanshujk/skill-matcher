package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync/atomic"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type progressMsg struct {
	downloaded int64
	total      int64
}

type finishedMsg struct {
	err error
}

var (
	globalDownloaded int64
	globalTotal      int64
	downloadDone     int32
	downloadError    error
)

type downloadModel struct {
	progress   progress.Model
	downloaded int64
	total      int64
	startTime  time.Time
	ctx        context.Context
	cancel     context.CancelFunc
	done       bool
	err        error
}

func newDownloadModel() downloadModel {
	ctx, cancel := context.WithCancel(context.Background())

	p := progress.New(
		progress.WithGradient("#00FFAA", "#0077FF"),
		progress.WithWidth(50),
	)

	// Reset global state
	atomic.StoreInt64(&globalDownloaded, 0)
	atomic.StoreInt64(&globalTotal, 0)
	atomic.StoreInt32(&downloadDone, 0)
	downloadError = nil

	return downloadModel{
		progress:  p,
		startTime: time.Now(),
		ctx:       ctx,
		cancel:    cancel,
	}
}

func (m downloadModel) Init() tea.Cmd {
	return tea.Batch(
		m.startDownload(),
		m.checkProgress(),
	)
}

func (m downloadModel) startDownload() tea.Cmd {
	return func() tea.Msg {
		go func() {
			err := downloadFile(m.ctx)
			downloadError = err
			atomic.StoreInt32(&downloadDone, 1)
		}()
		return nil
	}
}

func downloadFile(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, "GET", modelURL, nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(zipFile)
	if err != nil {
		return err
	}
	defer file.Close()

	total := resp.ContentLength
	atomic.StoreInt64(&globalTotal, total)

	buf := make([]byte, 32*1024) // 32KB chunks
	var downloaded int64

	for {
		select {
		case <-ctx.Done():
			os.Remove(zipFile)
			return fmt.Errorf("download cancelled")
		default:
		}

		n, err := resp.Body.Read(buf)
		if n > 0 {
			_, writeErr := file.Write(buf[:n])
			if writeErr != nil {
				return writeErr
			}
			downloaded += int64(n)
			atomic.StoreInt64(&globalDownloaded, downloaded)
		}

		if err == io.EOF {
			return nil // Success
		}

		if err != nil {
			return err
		}
	}
}

func (m downloadModel) checkProgress() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(time.Time) tea.Msg {
		downloaded := atomic.LoadInt64(&globalDownloaded)
		total := atomic.LoadInt64(&globalTotal)
		done := atomic.LoadInt32(&downloadDone) == 1

		if done {
			return finishedMsg{err: downloadError}
		}

		return progressMsg{
			downloaded: downloaded,
			total:      total,
		}
	})
}

func (m downloadModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.String() == "esc" {
			m.cancel()
			return m, tea.Quit
		}

	case progressMsg:
		m.downloaded = msg.downloaded
		m.total = msg.total

		var cmds []tea.Cmd

		if m.total > 0 {
			percent := float64(m.downloaded) / float64(m.total)
			progressCmd := m.progress.SetPercent(percent)
			cmds = append(cmds, progressCmd)
		}

		cmds = append(cmds, m.checkProgress())
		return m, tea.Batch(cmds...)

	case finishedMsg:
		m.done = true
		m.err = msg.err
		if msg.err == nil {
			m.downloaded = m.total
			cmd := m.progress.SetPercent(1.0)
			return m, tea.Sequence(cmd, tea.Quit)
		}
		return m, tea.Quit
	}

	// Always let the progress model handle any progress-related messages
	var cmd tea.Cmd
	progressModel, cmd := m.progress.Update(msg)
	m.progress = progressModel.(progress.Model)
	return m, cmd
}

func (m downloadModel) View() string {
	if m.err != nil {
		return fmt.Sprintf("\nDownload failed: %v\n", m.err)
	}

	if m.done && m.err == nil {
		return "\nDownload complete!\n"
	}

	percent := 0.0
	if m.total > 0 {
		percent = float64(m.downloaded) / float64(m.total)
	}

	elapsed := time.Since(m.startTime).Seconds()

	speed := 0.0
	if elapsed > 0 && m.downloaded > 0 {
		speed = float64(m.downloaded) / 1024 / 1024 / elapsed
	}

	remainingBytes := float64(m.total - m.downloaded)
	eta := 0.0
	if speed > 0 && m.total > 0 {
		eta = remainingBytes / (speed * 1024 * 1024)
	}

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FFAA"))

	infoStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#AAAAAA"))

	if m.total == 0 {
		return fmt.Sprintf(
			"\n%s\n\n%s\n\n%s\n\n(Press ESC to cancel)\n",
			titleStyle.Render("Downloading Embedding Model"),
			"Connecting...",
			infoStyle.Render("Starting download..."),
		)
	}

	return fmt.Sprintf(
		"\n%s\n\n%s\n\n%s\n%s\n%s\n\n(Press ESC to cancel)\n",
		titleStyle.Render("Downloading Embedding Model"),
		m.progress.View(),
		infoStyle.Render(fmt.Sprintf("%.0f%%", percent*100)),
		infoStyle.Render(fmt.Sprintf("%.2f MB / %.2f MB",
			float64(m.downloaded)/1024/1024,
			float64(m.total)/1024/1024)),
		infoStyle.Render(fmt.Sprintf("%.2f MB/s | %.0f sec remaining", speed, eta)),
	)
}

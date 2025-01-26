package bot

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/MattSilvaa/leethero/internal/config"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

type LeetHero struct {
	ctx    context.Context
	cancel context.CancelFunc
	config *config.Config
}

func New(cfg *config.Config) (*LeetHero, error) {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.NoFirstRun,
		chromedp.NoDefaultBrowserCheck,
	}

	if cfg.Headless {
		opts = append(opts, chromedp.Headless)
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, _ = chromedp.NewContext(ctx)

	return &LeetHero{
		ctx:    ctx,
		cancel: cancel,
		config: cfg,
	}, nil
}

func (h *LeetHero) setCookie(ctx context.Context) error {
	fmt.Println("Setting LeetCode session cookie...")
	return chromedp.Run(ctx,
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(14 * 24 * time.Hour))
			cookie := []*network.CookieParam{
				{
					Name:     "LEETCODE_SESSION",
					Value:    h.config.LeetCodeSession,
					Domain:   ".leetcode.com",
					Path:     "/",
					Expires:  &expr,
					Secure:   true,
					HTTPOnly: true,
				},
			}
			return network.SetCookies(cookie).Do(ctx)
		}),
	)
}

func (h *LeetHero) solveProblem(slug string) error {
	fmt.Printf("Solving problem: %s\n", slug)
	editorSelector := `.monaco-editor textarea.inputarea`
	submitButton := `[data-e2e-locator="console-submit-button"]`

	solutionCode := Solutions[slug]
	var isPython3 bool

	return chromedp.Run(h.ctx,
		chromedp.Navigate(fmt.Sprintf("https://leetcode.com/problems/%s/", slug)),
		chromedp.WaitVisible(editorSelector),
		chromedp.Evaluate(`
    document.evaluate('//div[contains(text(), "Python3")]', document, null, XPathResult.FIRST_ORDERED_NODE_TYPE, null).singleNodeValue !== null
`, &isPython3),
		chromedp.ActionFunc(func(ctx context.Context) error {
			if !isPython3 {
				if err := chromedp.Click(`//button[contains(text(), 'C++')]`).Do(ctx); err != nil {
					return err
				}
				if err := chromedp.WaitVisible(`//div[contains(text(), "Python3")]`).Do(ctx); err != nil {
					return err
				}
				if err := chromedp.Click(`//div[contains(text(), "Python3")]`).Do(ctx); err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.WaitVisible(editorSelector),
		chromedp.Focus(editorSelector),

		chromedp.Evaluate(`
           const editor = document.querySelector('.monaco-editor');
           const textarea = editor.querySelector('textarea');
           const model = editor.querySelector('.view-lines');

           // Clear input using multiple methods
           textarea.value = '';
           model.textContent = '';
           textarea.dispatchEvent(new Event('input', { bubbles: true }));
           textarea.dispatchEvent(new Event('change', { bubbles: true }));
        `, nil),

		chromedp.SendKeys(editorSelector, solutionCode),

		chromedp.WaitVisible(submitButton),
	)
}

func (h *LeetHero) Run() error {
	defer h.cancel()
	if err := h.setCookie(h.ctx); err != nil {
		return fmt.Errorf("failed to set cookie: %v", err)
	}

	for _, problem := range h.config.Problems {
		log.Printf("Solving problem: %s", problem)
		if err := h.solveProblem(problem); err != nil {
			log.Printf("Failed to solve %s: %v", problem, err)
			continue
		}
		log.Printf("Successfully solved: %s", problem)
	}

	return nil
}

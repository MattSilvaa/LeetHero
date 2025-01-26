# LeetHero 🦸‍♂️

Ever dreamed of a pristine, all-green submission history? Tired of manually grinding through LeetCode problems? **LeetHero** is here to save your day!

## Features 🚀
- Cookie-based authentication
  - IDK maybe I'll do auto-login eventually when I research how to actually bypass Cloudflare
- Lightning-fast problem submissions
- Pre-loaded solutions that actually work
- Your ticket to algorithmic glory

## Installation 🛠️
```bash
git clone https://github.com/yourusername/leethero
cd leethero
go install
```

## Usage 💻
1. Get your LEETCODE_SESSION cookie from browser:
    - Login to LeetCode
    - Open DevTools (F12)
    - Go to Application > Cookies
    - Copy LEETCODE_SESSION value


2. Run using environment variable:
```bash
export LEETCODE_SESSION="your_cookie_value"
./leethero
```

Or using command line flag:
```bash
./leethero -cookie="your_cookie_value"
```

Additional flags:
```bash
  -headless=false     # Show browser window
  -delay=5s          # Set delay between actions
  -problems="two-sum,add-two-numbers"  # Specify problems

3. Run LeetHero:
```bash
./leethero
```

Or with custom flags:
```bash
./leethero --headless=true --problems="two-sum,add-two-numbers"
```

## Warning ⚠️
With great power comes great responsibility. Use this tool wisely, or risk:
- Actually learning nothing
- Getting caught
- Eternal shame from real LeetCode gods
- Your future interviewer finding this repo

## Legal Stuff 🤓
This tool is for educational purposes only. Like Batman's gadgets, use it for good, not evil.

Remember: Heroes don't cheat. But they do automate. 😉

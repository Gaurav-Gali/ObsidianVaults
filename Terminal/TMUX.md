![Source](https://youtu.be/UxbiDtEXuxg)

**Note**: ==control+shift+b== is the prefix key that will activate the tmux commands

1. Installing tmux
```bash
brew install tmux
```

2. Entering tmux
```bash
tmux
```

3. Exiting tmux : prefix key then d

4. Creating a new session
```bash
tmux new -s session_name
```

5. Reattaching to the same tmux session
```bash
tmux attach -t session_name
```

6. Killing a session
```bash
tmux kill-ses -t session_name
```

7. Vertical split : prefix key then %
8. Horizontal split : prefix key then " (single double quotes)
9. Moving between panes : prefix key then press the arrow based on which pane to move
10. To kill a sessions : first be in that session and type exit.
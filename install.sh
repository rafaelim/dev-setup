#!/usr/bin/env bash
set -e

RIG_DIR="$HOME/.rig"
REPO="https://github.com/rafaelim/dev-setup"
BINARY_URL="https://github.com/rafaelim/dev-setup/releases/latest/download/rig"

echo ">> Bootstrapping rig..."

sudo apt-get update -y
sudo apt-get install -y curl git

# Clone or update repo
if [ -d "$RIG_DIR/.git" ]; then
    echo ">> Updating existing rig at $RIG_DIR..."
    git -C "$RIG_DIR" pull
else
    echo ">> Cloning rig to $RIG_DIR..."
    git clone "$REPO" "$RIG_DIR"
fi

# Download rig binary
echo ">> Downloading rig CLI..."
mkdir -p "$HOME/.local/bin"
curl -fsSL "$BINARY_URL" -o "$HOME/.local/bin/rig"
chmod +x "$HOME/.local/bin/rig"

# Ensure ~/.local/bin is in PATH
if ! echo "$PATH" | grep -q "$HOME/.local/bin"; then
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> "$HOME/.bashrc"
    export PATH="$HOME/.local/bin:$PATH"
fi

# Run installers
echo ">> Running installers..."
for script in "$RIG_DIR/runs/"*; do
    [ -x "$script" ] || continue
    echo "   -> $(basename "$script")"
    "$script"
done

# Run post-install commands
echo ">> Running post-install commands..."
for script in "$RIG_DIR/cmds/"*; do
    [ -x "$script" ] || continue
    echo "   -> $(basename "$script")"
    "$script"
done

# Sync dotfiles
echo ">> Syncing dotfiles..."
"$HOME/.local/bin/rig" sync

echo ""
echo ">> Rig ready. Reload your shell: source ~/.bashrc"
echo ">> Commands: rig sync | rig push | rig upgrade"

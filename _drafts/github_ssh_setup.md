echo Setting up ssh ls -al ~/.ssh

# Lists the files in your .ssh directory, if they exist

ssh-keygen -t ed25519 -C "vince@ory.sh"

# Start the ssh-agent in the background.

eval "$(ssh-agent -s)"

# check to see if your ~/.ssh/config file exists in the default location.

open ~/.ssh/config

# If the file doesn't exist, create the file.

touch ~/.ssh/config

# Open your ~/.ssh/config file, then modify the file, replacing ~/.ssh/id_ed25519 if you are not using the default location and name for your id_ed25519 key.

open ~/.ssh/config

# Add the following to config

`Host * AddKeysToAgent yes UseKeychain yes IdentityFile ~/.ssh/id_ed25519`

# Add your SSH private key to the ssh-agent and store your passphrase in the keychain. If you created your key with a different name, or if you are adding an existing key that has a different name, replace id_rsa in the command with the name of your private key file.

ssh-add -K ~/.ssh/id_ed25519

# Copy the SSH key to your clipboard.

# If your SSH key file has a different name than the example code, modify the filename to match your current setup. When copying your key, don't add any newlines or whitespace.

pbcopy < ~/.ssh/id_rsa.pub

(Add SSH Key to
Github)[https://docs.github.com/en/free-pro-team@latest/github/authenticating-to-github/connecting-to-github-with-ssh]

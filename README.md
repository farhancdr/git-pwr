# git-pwr
A tiny cli for your daily git usage


**git-pwr** is a handy command-line tool built in Go that empowers you to manage Git branches quickly and efficiently. It offers two essential commands to make your Git experience smoother:

**Key Features:**

- **copy-branch:**
   - Seamlessly list Git branches with the ability to navigate using up/down arrows.
   - Effortlessly copy a chosen branch name to your clipboard with a simple press of Enter.
   - Customize the number of branches displayed with the `-n` flag (e.g., `git-pwr copy-branch -n 20` to show 20 branches).

- **publish-branch:**
   - Create a new branch with ease by simply entering the desired branch name.
   - Publish the newly created branch to remote for seamless collaboration.

## Installation

```bash
bash -c "$(curl -fsSL https://raw.githubusercontent.com/farhancdr/git-pwr/main/install.sh)"
```
## Usage:

### copy-branch:

```bash
git-pwr copy-branch
git-pwr copy-branch -n <number_of_branches>
```

### publish-branch:

```bash
git-pwr publish-branch
```


## Contributing:

We welcome contributions! Feel free to submit pull requests or open issues for bug reports or feature requests.

License:

This project is licensed under the MIT License. See the LICENSE file for details.

**Stay productive and make branch management a breeze with git-pwr!**

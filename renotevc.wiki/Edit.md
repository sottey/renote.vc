Edit command could take "one" argument, and that has to be note-name. `renotevc edit [name]` <br> But also, it must not take any argument, so it'll show the full list of notes and ask you to select the note that you wanna edit. `renotevc edit` <br>

And then it will open the file with the defined editor (by you) to edit it.

- **Usage**: `renotevc edit [flags]`
- **Aliases**: `edit`, `overwrite`, `update`
- **Flags**: `-h`, `--help` -> help for edit

### Edit (take node from arguments)
<img width="600" alt="edit" src="https://user-images.githubusercontent.com/59066341/148656096-8b4d06e6-1fd7-413d-aabb-07e39e1c5e16.png">

### Edit (choose node from suggestions)
<img width="600" alt="edit" src="https://user-images.githubusercontent.com/59066341/148656083-511f05f8-79d1-40dd-b2f1-f8cff7433d2a.png">
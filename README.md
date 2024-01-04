# typora_image_uploader

<div align="center">
  <img src="assets/icon.jpg" alt="icon" width="400" />
</div>


## intro

auto upload image in Typora to github repository

![preview](assets/preview.gif)



## usage

1. create a repository in github

2. [Creating a github personal access token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/managing-your-personal-access-tokens#creating-a-personal-access-token-classic)

3. download from [release](https://github.com/obgnail/typora_image_uploader/releases/latest)

4. set `Custom Command` in Typora

   ```sh
   your/path/to/imageUploader_windows_amd64.exe -token=your_token -owner=your_github_name -repo=your_repo_name
   ```



![setCustomUploader](assets/setCustomUploader.png)



![test](assets/test.png)


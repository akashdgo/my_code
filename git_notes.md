## git steps 

```
[akashdas@akash-mac github.com]$ cd my_code/
[akashdas@akash-mac my_code (master)]$ git init
[akashdas@akash-mac my_code (master)]$ cd add --all
[akashdas@akash-mac my_code (master)]$ git commit -m "first commit"
[akashdas@akash-mac my_code (master)]$ git remote add origin https://github.com/akashdgo/my_code.git
[akashdas@akash-mac my_code (master)]$ git push -u origin my_code
```

### if the file test.md is deleted from the web

```
[akashdas@akash-mac my_code (master)]$ git push -u origin master
To https://github.com/akashdgo/my_code.git
 ! [rejected]        master -> master (fetch first)
error: failed to push some refs to 'https://github.com/akashdgo/my_code.git'
hint: Updates were rejected because the remote contains work that you do
hint: not have locally. This is usually caused by another repository pushing
hint: to the same ref. You may want to first integrate the remote changes
hint: (e.g., 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.

[akashdas@akash-mac my_code (master)]$ git pull --all
Fetching origin
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (2/2), done.
remote: Total 3 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (3/3), done.
From https://github.com/akashdgo/my_code
   1da3fbd..cd43f05  master     -> origin/master
There is no tracking information for the current branch.
Please specify which branch you want to merge with.
See git-pull(1) for details.

    git pull <remote> <branch>

If you wish to set tracking information for this branch you can do so with:

    git branch --set-upstream-to=origin/<branch> master


[akashdas@akash-mac my_code (master)]$ git push origin master -f
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 4 threads.
Compressing objects: 100% (3/3), done.
Writing objects: 100% (4/4), 421 bytes | 421.00 KiB/s, done.
Total 4 (delta 2), reused 0 (delta 0)
remote: Resolving deltas: 100% (2/2), completed with 1 local object.
To https://github.com/akashdgo/my_code.git
 + cd43f05...be21359 master -> master (forced update)
```

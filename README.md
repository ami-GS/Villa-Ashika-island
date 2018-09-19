[![Build Status](https://travis-ci.org/ami-GS/Villa-Ashika-island.svg?branch=master)](https://travis-ci.org/ami-GS/Villa-Ashika-island)


# Villa-Ashika-island
web for villa Ashika island

## Prerequisit
- Hugo
- Theme
  - clone theme and copy the directory to `./themes/` (mkdir `./themes` at first)

detailed installation can be seen in `.travis.yml`

## Create New Post
```
hugo new content/post/title.md
```
and edit headers and write down contents

## Test

Under this directory
```
>> hugo server
```

## Deploy
Just push change to master, travis CI will deploy S3 as origin server.

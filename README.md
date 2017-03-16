# pacnews

Arch news reader and pacman interrupter when there is news

This application offers a way to view the Arch news and prevents you from updating/installing/removing things until you read that news item.

## Usage

```shell
pacnews c # checks if there are new news items, and if so, returns exit code 1
```

```shell
pacnews r # lists all the news items you haven't seen before and marks them as read
```

## Workings

* It stores a bolt database with all the read items in: ```/var/cache/pacnews.db```
* It installs a hook in ```/etc/pacman.d/hooks/pacnews.hook```

#!/bin/sh

cmd="sass $@"

$cmd web/styles/styles.scss:../public/css/styles.css &
$cmd web/styles/home.scss:../public/css/home.css &
$cmd web/styles/post/posts.scss:../public/css/posts.css &
$cmd web/styles/post/post.scss:../public/css/post.css &


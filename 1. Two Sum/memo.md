## step1
最初は総当たりで書こうと思ったが、ペアの値をmapから引く方が個人的に直感的だと感じこの方法で実装した。（コードではなく人力でやる場合を想像した時に飲み込みやすかった）

## step2
mapの存在確認はokからより意図が伝わりやすいexistsにした。また、見つからなかった場合の戻り値はより一般的なnilにした。https://go.dev/wiki/CodeReviewComments
> the nil slice is the preferred style.

## step3
step2の方法で3回実装

## step3.2
PRを作成した後に見つけた他の方へのコメント
https://github.com/officermushroom-dotcom/leetcode_practice/pull/1/changes#r4306518534
確かにcomplementの方が意図がぶれにくい
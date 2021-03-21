# nippo
日報を手軽に記録しよう．

## なにを記録する？
|記録事項|説明|
|----|----|
|作業日||
|今日の作業||
|完了した作業||
|明日の作業||

## Database構造
### `nippo`
```sql
create table nippo.nippo (
    id int not null primary key auto_increment,
    date varchar(20),
    content varchar(5000)
);
```

### `task`
```sql
create table nippo.task (
    id int not null primary key auto_increment,
    title varchar(50),
    content varchar(1000),
    created_date varchar(20),
    deadline_date varchar(20),
    done_date varchar(20)
);
```

## References
- [goでhttpサーバ起動と同時にブラウザを開く例](https://gist.github.com/niratama/6b0117c6c6f2d21b5687)
- [Goでwebアプリを開発してみよう](https://www.slideshare.net/takuyaueda967/goweb-69949279)
- [[Golang] ファイル読み込みサンプル](https://qiita.com/tchnkmr/items/b686adc4a7e144d48755)
- [GolangでHTMLのtemplateに値を渡す方法](https://qiita.com/tetsuzawa/items/0d043ad76b9705cdbb79)
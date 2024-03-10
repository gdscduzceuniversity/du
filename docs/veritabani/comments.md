# Comments collection'u
Bu koleksiyonda yorumlarla alakalı her türlü bilgiler yer alacaktır.

## _id
id alanı tipi objectid olup mongo'nun bize ön tanımlı vermiş olduğu numaralandırma için otomatik olarak kullanılmaktadır.

## author
bu alan string tipinde olup yorumun yazarının id'sini barındırır.

## content
bu alan string tipinde olup yorumun içeriğini barındırır.

## likes
bu alan array tipinde olup yorumu beğenen kullanıcıların id'lerini barındırır.

## createdAt
bu alan tarih tipinde olup yorumun oluşturulma tarihini barındırır.

## isDeleted
bu alan bool tipinde olup yorumun silinip silinmediğini barındırır. 1: silindi, 0: silinmedi
```
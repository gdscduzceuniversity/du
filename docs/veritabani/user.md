# Users collection'u
Bu koleksiyonda projede kullanıcıyla alakalı her türlü bilgiler yer alacaktır.

## _id
id alan'ı, tipi nesne numarası şeklinde mongo'nun bize ön tanımlı vermiş olduğu numaralandırma için otomatik olarak kullanılmaktadır.

## email
bu alan string tipinde olup kullanıcının email'ini barındırır unique olması mühimdir.

# firstName
bu alan string tipinde olup kullanıcının adını barındırır, bu alan değiştirilebilir.

# lastName
bu alan string tipinde olup kullanıcının soyadını barındırır, bu alan değiştirilebilir.

# passHash
bu alan string tipinde olup kullanıcının parolasının hash'ini tutar, bu alan değiştirilebilir.

# slug
bu alan string tipinde olup kullanıcının sayfasına giden url path'ini verir, bu alan değiştirilebilir lakin unique'dir.

# bio
bu alan string tipinde olup kullanıcnın hakkında açıklaması yer alır bu alan değiştirilebilir.

# iPath
bu alan string tipinde olup resmin dosya sisteminde ki adresi verilmiştir.

# isActive
bu alan bool tipinde olup kullanıcnın inaktif ya da aktif olduğu durumunun bilgisini içerir. 1: aktif, 0: inaktif

# createdAt
bu alan tarih tipinde olup kullanıcının oluşturulma tarihini barındırır.

# communities
bu alan array tipinde olup kullanıcının hangi topluluklara üye olduğunu id tipinde bir array olarak barındırır.

# isDeleted
bu alan bool tipinde olup kullanıcının silinip silinmediğini barındırır. 1: silindi, 0: silinmedi
```


# crud

全是get请求
``` 
# create
# mutation mutation_create_name2 {
#   create(name: "Inca Kola", info: "Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)", price: 1.99) {
#     id
#     name
#     info
#     price
#   }
# }

# search one
# query{
#   product(id:1){
#     id
#     name
#     price
#   }
# }

# search many
# need alias
# query{
#  p1: product(id:1){
#     id
#     name
#     price
#   }
#  p2: product(id:2){
#     name
#     price
#   }
# }

# search all
# query{
#   list{
#     id
#     name
#     price
#   }
# }

# update
# mutation{
#   update(id:1,price:3.95){
#     id
#     name
#     info
#     price
#   }
# }

# delete
# mutation{
#   delete(id:1){
#     info
#     price
#   }
# }
```

`

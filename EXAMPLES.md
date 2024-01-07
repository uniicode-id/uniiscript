### Hello World
```unii
print("Hello World");
```

### Comment
#### Single Line Comment
```unii
// Single Line Comment
```
#### Multi Line Comment
```unii
/*
Multi 
Line
Comment
*/
```

### Variable
```unii
var name = "John";
print(name);

name = "Andre";
print(name);
```

### Const
```unii
const Pi = 3.14;
print(Pi);

Pi = 3.141; // ERROR!!!
```
Tidak ada immutable variable, maaf

### Data Type
```
var name: Str = "Andre";
var age: Int = 19;
var height: Float = 178.9;
var isMarried: Bool = false;
var bloodType: Char = 'O';
var girlFriend: String? = null;
```
Tidak ada `undefined`, yeay!

### Array
```
var skills: Str[] = ["Coding", "Cooking", "Gaming"];
print(skills); // ["Coding", "Cooking", "Gaming"]

skills.add("Writing");
print(skills); // ["Coding", "Cooking", "Gaming", "Writing"]

skills.remove(1);
print(skills); // ["Coding", "Gaming", "Writing"]

print(skills.length); // 3
```

### Operator
#### Arithmetic
```unii
print(3 + 3); // penjumlahan
print(3 - 3); // pengurangan
print(3 * 3); // perkalian
print(3 / 3); // pembagian
print(3 % 2); // modulus (sisa bagi)
print(3 ** 2); // perpangkatan
print(3 // 2); // pembagian tanpa koma
```
#### Comparison
```unii
print(3 == 3); // apakah sama
print(3 != 3); // apakah tidak sama
print(3 > 3); // apakah lebih besar
print(3 >= 3); // apakah lebih besar atau sama
print(3 < 3); // apakah lebih kecil
print(3 <= 3); // apakah lebih kecil atau sama
```
#### Logical
```unii
print(!true); // negasi
print(true && false); // and
print(true || false); // or
```
#### Precedence dan Grouping
```unii
print(9 + 3 * (2 + 5));
```

### Control Flow
#### If
```unii
var age = 18;
if (age >= 18) {
    print("Masuk");
}
```
#### If Else
```unii
var age = 18;
if (age >= 18) {
    print("Masuk");
} else {
    print("Tidak Masuk");
}
```
#### If Else-If Else
```unii
var age = 18;
if (age >= 18) {
    print("Masuk");
} else if (age >= 60) {
    print("Sehat - Sehat");
} else {
    print("Tidak Masuk");
}
```
#### Switch Case
```unii
var key = 'W';
switch (key) {
    case 'A':
        print("Kiri");
    case 'S':
        print("Mundur");
    case 'D':
        print("Kanan");
    case 'W':
        print("Maju");
    default:
        print("Diam");
}
```
#### While Loop
```unii
var i = 0;
while (i < 10) {
    print(i);
    i++;
}
```
Tidak ada do-while, maaf
#### For Loop
```
for (var i = 0; i < 10; i++) {
    print(i);
}
```

### Function
#### Basic
```unii
fun eat() {
    print("I'm Eating!");
}

eat();
```
#### Parameter
```unii
fun eat(food: Str) {
    print("I'm Eating $food");
}
```
#### Return
```unii
fun eat(food: Str) {
    return "I'm Eating $food";
}

fun drink(drinks: Str): Str {
    return "I'm Drinking $drinks";
}
```
Hanya bisa mengembalikan satu nilai, maaf
#### Default Parameter
```unii
fun eat(food: Str = "Rice") {
    print("I'm Eating $food");
}
```
Parameter yang default harus ditaruh setelah parameter non-default
#### Rest Parameter
```unii
fun printNumbers(...numbers: Int[]) {
    foreach (number in numbers) {
        print(number);
    }
}
```
Parameter yang rest harus ditaruh di paling akhir
#### First Class Function
```unii
var say = fun(message: String) {
    print("Hi $message");
}

var check: Fun<Bool, Str> = fun(data: Str): Bool {
    return data != "";
}

if (check("foo")) {
    say("Success!");
}
```

### OOP
#### Class, Property, Method, Constructor, Modifier, dan Instance
```unii
class User {
    var name: Str;
    var email: Str;
    var password: Str;
    
    constructor(name: Str, email: Str, password: Str) {
        this.name = name;
        this.email = email;
        this.password = password;
    }
    
    fun register() {
        debug("Register");
    
        db.users.create(
            name,
            email,
            password
        );
    }
    
    private fun debug(info: Str) {
        print("$info: $this");
    }
}

var user1 = User(
    "Andre",
    "andreee@gmail.com",
    "urangSunda"
);
user1.register();
```
#### Enum
```unii
enum AdminAccess {
    CREATE,
    READ,
    UPDATE,
    DELETE
}
```
#### Inheritance
```unii
class Admin extends User {
    private var accesses: AdminAccess[];

    constructor(
        name: Str,
        email: Str,
        password: Str,
        accesses: AdminAccess[]
    ) {
        super(name, email, password);
        this.accesses = accesses;
    }
}
```
#### Encapsulation
```unii
class Admin extends User {
    private var accesses: AdminAccess[];

    constructor(
        name: Str,
        email: Str,
        password: Str,
        accesses: AdminAccess[]
    ) {
        super(name, email, password);
        this.accesses = accesses;
    }
    
    fun getAccesses() {
        return accesses;
    }
    
    
    // fun setAccesses(accesses: AdminAccess[]) {
    //     this.accesses = accesses;
    // }
}
```

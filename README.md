## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

## API Endpoints and Usage

### Authorization Endpoints

##### Create user
```
POST /auth/user

{
    "email":     "test@email.com",
    "password":  "password"
}
```
```
Returns:
{
    "email":     "test@email.com",
    "token":     "iOiJIUzI1NiIsInR5cCI6IkpX",
    "userId":    "1234",
    "monsters":  null
}
```

##### Login
```
POST /auth/login

{
    "email":     "test@email.com",
    "password":  "password"
}
```
```
Returns:
{
    "email":     "test@email.com",
    "token":     "iOiJIUzI1NiIsInR5cCI6IkpX",
    "userId":    "1234",
    "monsters":  [ ... ]
}
```

### User Endpoints

##### Get user data by ID
```
GET /user/{userID}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "email":     "test@email.com",
    "token":     "iOiJIUzI1NiIsInR5cCI6IkpX",
    "id":        "1234",
    "monsters":  [
        {
            "ID":       "2345",
            "No":       66,
            "Name":     "Machop",
            "Type":     "Fighting",
            "Hp":       70,
            "Attack":   80,
            "Defense":  50,
            "Attacks":  [
                {
                    "SlotNo":       0,
                    "MonsterNo":    66,
                    "Name":         "Karate Chop",
                    "Type":         "Normal",
                    "Power":        50,
                    "Accuracy":     100,
                    "AnimationID":  2
                },
                { ... }
            ]
        },
        { ... }
    ]
}
```

##### Get user animations by ID
```
GET /user/{userID}/animations

header: "authorization":  <auth_token>
```
```
Returns:
{
    "animations": [
        {
            "MonsterNo":    66,
            "AnimationIDs": [2, 3, 4, 5]
        },
        {
            "MonsterNo":    7,
            "AnimationIDs": [2, 6, 8, 9]
        }
    ]
}
```

##### Add monster to user
```
POST /user/monster/

{
    "monsterNo":  66
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "monster added",
}
```

##### Rename monster
```
PUT /user/monster/

{
    "monsterID":  "123",
    "name":       "MaCHOMP"
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "monster renamed to MaCHOMP",
}
```

##### Add attack to user's monster
```
POST /user/monster/attack/

{
    "attackID":   "234",
    "monsterID":  "123",
    "slotNo":     2
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "attack added",
}
```

### Dex Endpoints

##### Create monster
```
POST /dex/monster

{
    "monsterNo":  66,
    "name":       "Machop",
    "type":       "Fighting",
    "hp":         70,
    "attack":     80,
    "defense":    50,
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "monster": {
        "ID":       "",
        "No":       66,
        "Name":     "Machop",
        "Type":     "Fighting",
        "Hp":       70,
        "Attack":   80,
        "Defense":  50,
        "Attacks":  null
    }
}
```

##### Create attack
```
POST /dex/attack

{
    "monsterNo":    66,
    "name":         "Karate Chop",
    "type":         "Normal",
    "power":        50,
    "accuracy":     100,
    "animationID":  2
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "attack": {
        "SlotNo":       0,
        "MonsterNo":    66,
        "Name":         "Karate Chop",
        "Type":         "Normal",
        "Power":        50,
        "Accuracy":     100,
        "AnimationID":  2
    }
}
```

## Database Schema

### User
| Field        | Type      | Description
| ------------ | :-------: | ---------
| ID           | string    | --
| AuthToken    | string    | --
| Email        | string    | --
| Password     | string    | --
| PasswordHash | string    | --
| Monsters	   | []Monster | Array of user's monsters

### Monster
> Note: Monsters exist in the `dex` database under the `monsters` collection. These monsters have all base stats with no set ID, and are identified by the `No` field.

| Field        | Type      | Description
| ------------ | :-------: | ---------
| ID           | string    | Set once added to a user
| No           | int       | Monster number in Dex
| Name         | string    | --
| Type         | string    | --
| Hp           | int       | --
| Attack	   | int       | --
| Defense      | int       | --
| Attacks      | int       | Array of monster's learned attacks

### Attack
> Note: Attacks are located in the `dex` database under the `attacks` collection. 

| Field        | Type      | Description
| ------------ | :-------: | ---------
| SlotNo       | int       | Position in monster's attack array, value of 0-3.
| Name         | string    | --
| Type         | string    | --
| Power        | int       | --
| Accuracy	   | int       | --
| AnimationID  | int       | Corresponds to animation in .dae

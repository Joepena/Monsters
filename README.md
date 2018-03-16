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
    "email":   "test@email.com",
    "token":   "iOiJIUzI1NiIsInR5cCI6IkpX",
    "userId":  "1234"
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
    "token":   "iOiJIUzI1NiIsInR5cCI6IkpX",
    "userId":  "1234"
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
            ],
            "hits": 230,
            "misses": 122,
            "damageDealt": 14000,
            "damageReceived": 2000
        },
        { ... }
    ],
    "battles": [
        {
            "victorID": "1234",
            "loserID": "5678",
            "date": "2018-02-25T21:06:59.222-05:00",
            "location": {
                "x": 23.389,
                "y": -24.1342
            }
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

##### Update monster stats
```
PUT /user/monster/stats

{
	"monsterID": "123",
	"hits": 2,
	"misses": 4,
	"damageDealt": 200,
	"damageReceived": 150
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "monster stats updated"
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

##### Add battle results to user
```
POST /user/battle

{
	"victorID": "1234",
	"loserID": "5678",
	"location": {
		"x": 23.3890,
		"y": -24.1342
	}
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "battle added to users 1234 and 5678",
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
        "Attacks":  null,
        "hits": 0,
        "misses": 0,
        "damageDealt": 0,
        "damageReceived": 0
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
| Battles	   | []Battle  | Array of battles user has fought

### Monster
> Note: Monsters exist in the `dex` database under the `monsters` collection. These monsters have all base stats with no set ID, and are identified by the `No` field.

| Field          | Type      | Description
| -------------- | :-------: | ---------
| ID             | string    | Set once added to a user
| No             | int       | Monster number in Dex
| Name           | string    | --
| Type           | string    | --
| Hp             | int       | --
| Attack	     | int       | --
| Defense        | int       | --
| Attacks        | []Attack  | Array of monster's learned attacks
| Hits           | int       | Total number of successful attacks
| Misses         | int       | Total number of missed attacks
| DamageDealt    | int       | Total damage done to other monsters
| DamageReceived | int       | Total damage received from other monsters

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

### Battle

| Field        | Type      | Description
| ------------ | :-------: | ---------
| VictorID     | string    | UserID of the victor
| LoserID      | string    | UserID of the loser
| Date         | time.Time | Date and time set by server
| Location     | Location  | Geospatial coordinates

### Location

| Field        | Type      | Description
| ------------ | :-------: | ---------
| X            | float32   | Longitude
| Y            | float32   | Latitude

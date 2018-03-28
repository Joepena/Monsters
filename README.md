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
    "id": "1234",
    "token": "iOiJIUzI1NiIsInR5cCI6IkpX",
    "email": "test@email.com",
    "monsters": [],
    "battleStats": {
        "wins": 0,
        "loses": 0
    }
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
    "id":    "1234",
    "token":     "iOiJIUzI1NiIsInR5cCI6IkpX",
    "email":     "test@email.com",
    "monsters":  [ ... ],
    "battleStats": {
        "wins":   40,
        "losses": 12
    }
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
    "email":  "test@email.com",
    "id":     "1234",
    "battleStats": {
        "wins":    40,
        "losses":  12
    },
    "monsters":  [
        {
            "monsterID":  "2345",
            "monsterNo":  66,
            "assetIDSet": {
                "texture1ID": 1,
                "texture2ID": 2,
                "ios":{
                    "daeID": 3,
                    "animationSet": [
                        {
                          "name": "attack",
                          "assetID": 4
                        },
                        {
                          "name": "faint",
                          "assetID": 5
                        },
                        {
                          "name": "hit",
                          "assetID": 6
                        },
                        {
                          "name": "intro",
                          "assetID": 7
                        },
                        {
                          "name": "standing",
                          "assetID": 8
                        }
                      ]
                    },
                "android": {
                    "blendID": 9
                }
            },
            "name":       "Machop",
            "type":       "Fighting",
            "hp":         70,
            "attack":     80,
            "defense":    50,
            "attacks":    [
                {
                    "slotNo":       0,
                    "monsterNo":    66,
                    "name":         "Karate Chop",
                    "type":         "Normal",
                    "power":        50,
                    "accuracy":     100,
                    "animationID":  2
                },
                { ... }
            ],
            "stats": {
                "hits":             230,
                "misses":           122,
                "damageDealt":      14000,
                "damageReceived":   2000,
                "enemiesFought":    42,
                "enemiesDefeated":  12,
                "faints":           4
            }
        },
        { ... }
    ]
}
```

##### Get a user's assetSet by ID
```
GET /user/{userID}/assets

header: "authorization":  <auth_token>
```
```
Returns:
{
    "assets": [
        {
            "Name": "Squirtle",
            "MonsterNo": 7,
            "AssetSet": {
                "texture1ID": 1,
                "texture2ID": 2,
                "ios": {
                    "daeID": 3,
                    "animationSet": [
                        {
                            "name": "attack",
                            "assetID": 4
                        },
                        {
                            "name": "faint",
                            "assetID": 5
                        },
                        {
                            "name": "hit",
                            "assetID": 6
                        },
                        {
                            "name": "intro",
                            "assetID": 7
                        },
                        {
                            "name": "standing",
                            "assetID": 8
                        }
                    ]
                },
                "android": {
                    "blendID": 9
                }
            }
        },
        .....
    ],
    "userID": "1"
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
	"stats": {
        "hits":             2,
        "misses":           4,
        "damageDealt":      200,
        "damageReceived":   150,
        "enemiesFought":    1,
        "enemiesDefeated":  1,
        "faints":           0
	}
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

##### Add battle result to user
```
POST /user/battle

{
    "wins":   1,
    "losses": 0
}

header: "authorization":  <auth_token>
```
```
Returns:
{
    "status": "battle stats updated"
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
        "monsterID":  "",
        "monsterNo":  66,
        "name":       "Machop",
        "type":       "Fighting",
        "hp":         70,
        "attack":     80,
        "defense":    50,
        "attacks":    null,
        "stats": {
            "hits":             0,
            "misses":           0,
            "damageDealt":      0,
            "damageReceived":   0,
            "enemiesFought":    0,
            "enemiesDefeated":  0,
            "faints":           0
        }
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
        "slotNo":       0,
        "monsterNo":    66,
        "name":         "Karate Chop",
        "type":         "Normal",
        "power":        50,
        "accuracy":     100,
        "animationID":  2
    }
}
```

## Database Schema

### User

| Field        | Type        | Description
| ------------ | :---------: | ---------
| ID           | string      | --
| AuthToken    | string      | --
| Email        | string      | --
| Password     | string      | --
| PasswordHash | string      | --
| Monsters	   | []Monster   | Array of user's monsters
| BattleStats  | BattleStats | Stores user's wins & losses


### BattleStats

| Field        | Type     | Description
| ------------ | :------: | ---------
| Wins         | int      | --
| Losses       | int      | --


### Monster
> Note: Monsters exist in the `dex` database under the `monsters` collection. These monsters have all base stats with no set ID, and are identified by the `No` field.

| Field           | Type      | Description
| --------------- | :-------: | ---------
| ID              | string    | Set once added to a user
| No              | int       | Monster number in Dex
| Assets          | AssetIDSet| Set of asset files that correspond to the monster
| Name            | string    | --
| Type            | string    | --
| Hp              | int       | --
| Attack	      | int       | --
| Defense         | int       | --
| Attacks         | []Attack  | Array of monster's learned attacks
| Stats           | Stats     | Monster's battle stats

### AssetIDSet

| Field           | Type      | Description
| --------------- | :-------: | ---------
| Texture1ID      | int       | --
| Texture2ID      | int       | --
| IOS             | struct    | --
| Android         | struct    | --


### Stats

| Field           | Type      | Description
| --------------- | :-------: | ---------
| Hits            | int       | Total number of successful attacks
| Misses          | int       | Total number of missed attacks
| DamageDealt     | int       | Total damage done to other monsters
| DamageReceived  | int       | Total damage received from other monsters
| EnemiesFought   | int       | Number of enemies fought
| EnemiesDefeated | int       | Number of enemies defeated
| Faints          | int       | Number of times defeated by a monster


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

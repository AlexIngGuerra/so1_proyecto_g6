const Paises = {
    "Paises": ["Argentina-Mexico",
        "España-Alemania",
        "Brasil-Argentina"]
}

const Datos = {
    "Argentina-Mexico,1": {
        "Pais": "Argentina-Mexico,1",
        "Predics": [
            {
                "Punteo": "5-2",
                "Votos": 3
            },
            {
                "Punteo": "2-3",
                "Votos": 1
            },
            {
                "Punteo": "1-4",
                "Votos": 15
            }
        ]
    },
    "Argentina-Mexico,2": {
        "Pais": "Argentina-Mexico,2",
        "Predics": [
            {
                "Punteo": "1-2",
                "Votos": 54
            },
            {
                "Punteo": "3-5",
                "Votos": 51
            },
            {
                "Punteo": "3-6",
                "Votos": 11
            }
        ]
    },
    "Argentina-Mexico,3": {
        "Pais": "Argentina-Mexico,3",
        "Predics": [
            {
                "Punteo": "2-2",
                "Votos": 5
            },
            {
                "Punteo": "3-5",
                "Votos": 8
            },
            {
                "Punteo": "2-7",
                "Votos": 5
            }
        ]
    },
    "Argentina-Mexico,4": {
        "Pais": "Argentina-Mexico,4",
        "Predics": [
            {
                "Punteo": "15-4",
                "Votos": 9
            },
            {
                "Punteo": "1-2",
                "Votos": 9
            },
            {
                "Punteo": "7-2",
                "Votos": 1
            }
        ]
    },

    "España-Alemania,1": {
        "Pais": "España-Alemania,1",
        "Predics": [
            {
                "Punteo": "5-2",
                "Votos": 3
            },
            {
                "Punteo": "2-3",
                "Votos": 1
            },
            {
                "Punteo": "1-4",
                "Votos": 15
            }
        ]
    },
    "España-Alemania,2": {
        "Pais": "España-Alemania,2",
        "Predics": [
            {
                "Punteo": "1-2",
                "Votos": 54
            },
            {
                "Punteo": "3-5",
                "Votos": 51
            },
            {
                "Punteo": "3-6",
                "Votos": 11
            }
        ]
    },
    "España-Alemania,3": {
        "Pais": "España-Alemania,3",
        "Predics": [
            {
                "Punteo": "2-2",
                "Votos": 5
            },
            {
                "Punteo": "3-5",
                "Votos": 8
            },
            {
                "Punteo": "2-7",
                "Votos": 5
            }
        ]
    },
    "España-Alemania,4": {
        "Pais": "España-Alemania,4",
        "Predics": [
            {
                "Punteo": "15-4",
                "Votos": 9
            },
            {
                "Punteo": "1-2",
                "Votos": 9
            },
            {
                "Punteo": "7-2",
                "Votos": 1
            }
        ]
    }
}

const Api_URL = "http://192.168.197.130:8000"

export async function getPaises() {
    try {
        const response = await fetch(Api_URL + '/GetPaises',
        {
                //const response = await fetch('/Hijos',{
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({})
            });

        const data = await response.json();
        console.log(data)
        if (data.Codigo == 200) {
            return data;
        }
        return null
    } catch (error) {
        console.error(error);
        return null
    }

}

export async function getData(Pais, Fase) {
    console.log(JSON.stringify({"Pais": Pais, "Fase": parseInt(Fase)}))
    try {
        const response = await fetch(Api_URL + '/GetPaisFase',
        {
                //const response = await fetch('/Hijos',{
                method: 'POST',
                headers: {
                    'Accept': 'application/json',
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({"Pais": Pais, "Fase": parseInt(Fase)})
            });

        const data = await response.json();
        console.log(data)
        if (data.Codigo == 200) {
            return data;
        }
        return null
    } catch (error) {
        console.error(error);
        return null
    }
}
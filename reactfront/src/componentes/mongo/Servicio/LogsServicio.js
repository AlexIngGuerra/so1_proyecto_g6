const Api_URL = "http://34.135.93.12.nip.io:8000"

export async function getLogs(){
    try {
        const response = await fetch(Api_URL+'/obtenerLogs',
        {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
        });
        const data = await response.json();
        console.log(data);
        if (data.Codigo == 200) {
            return data
        }
        return null;
    } catch (error) {
        console.log(error)
    }
}

export async function deleteLogs(){
    try {
        const response = await fetch(Api_URL+'/deleteLogs',
        {
            method: 'GET',
            headers: {
                'Accept': 'application/json',
                'Content-Type': 'application/json'
            },
        });
        return null;
    } catch (error) {
        console.log(error);
    }
}
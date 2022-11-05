import { Box, Button, Grid, GridItem, Heading, Select, Table, TableContainer, Text } from '@chakra-ui/react';
import {React, useState, useEffect} from 'react';
import * as API from './Servicio/LiveServicio';

import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend,
  } from 'chart.js';

import { Bar } from 'react-chartjs-2';

ChartJS.register(
    CategoryScale,
    LinearScale,
    BarElement,
    Title,
    Tooltip,
    Legend
  );

  export const options = {
    indexAxis: 'y',
    elements: {
      bar: {
        borderWidth: 1,
      },
    },
    responsive: true,
    plugins: {
      legend: {
        position: 'right',
      },
      title: {
        display: true,
        text: '',
      },
    },
  };

export default function Live() {
    const [Pais, setPais] = useState();
    const [Fase, setFase] = useState(1);

    const [Paises, setPaises] = useState([]);
    const [Datos, setDatos] = useState(null);

    const [DataPunteo, setDataPunteo] = useState([]);
    const [DataVotos, setDataVotos] = useState([]);
    
    const [Ver, setVer] = useState(false);

    const [Buscar, setBuscar] = useState(true);

    

    useEffect(() => {
        //setPaises(API.getPaises());
        API.getPaises().then(setPaises).catch(console.log);
        //setDatos(API.getData(Pais+','+Fase));
        API.getData(Pais, Fase).then(setDatos).catch(console.log);
        
        setDataPunteo([]);
        setDataVotos([]);

        Actualizar_Grafo();
        
    }, [Fase, Pais, Ver, Buscar]);

    const Actualizar_Grafo = () => {
        var auxPunt = [];
        var auxVotos = [];
        try{
            for (var i = 0; i<Datos.Fase.Predics.length; i++){
                auxPunt.push(Datos.Fase.Predics[i].Punteo);
                console.log(Datos.Fase.Predics[i].Punteo);
                auxVotos.push(Datos.Fase.Predics[i].Votos)
            }
            console.log("DATA A GRAFICAR");
            console.log(auxPunt);
            console.log(auxVotos);
            setDataPunteo(auxPunt)
            setDataVotos(auxVotos);
            labels = DataPunteo;
            data = {
                labels : labels,
                datasets: [{
                    label: 'Resultados',
                    data: DataVotos,
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.5)'
                }]
            }
        }catch(error){
            console.log('un error')
            console.log(error)
        };
        setVer(Ver)
    }

    var labels = DataPunteo;
    var data = {
        labels : labels,
        datasets: [{
            label: 'Resultados',
            data: DataVotos,
            borderColor: 'rgb(255, 99, 132)',
            backgroundColor: 'rgba(255, 99, 132, 0.5)'
        }]
    };

    var alto = window.innerHeight;

    const CambioPais = event => {
        setPais(event.target.value);
    }

    const CambioFase = event => {
        setFase(event.target.value);
    }
    return (
        <>
            <Grid h={alto} templateRows='repeat(1)'
                templateColumns='repeat(9)' gap={1}>
                
                <GridItem colStart={1} colEnd={10} rowStart={1} rowEnd={1}>
                    <Grid h='full' templateRows='repeat(9, fr)' templateColumns='repeat(17)' gap={1}>
                        <GridItem rowStart={0} rowEnd={1} >
                            <Grid h='full'>
                               <Button onClick={() => setBuscar(!Buscar)} bg='blue' color='white'>
                                Buscar
                               </Button>
                            </Grid>
                        </GridItem>
                        <GridItem rowStart={1} rowEnd={2} >
                            <Grid templateColumns='repeat(16)' templateRows='repeat(3)' h='full'>
                                <GridItem colStart={0} colEnd={1} rowStart={1} rowEnd={3}>

                                </GridItem>
                                <GridItem colStart={1} colEnd={8} h='full' rowStart={1} rowEnd={1}>
                                    <Box  margin='auto' marginBottom='-10%' h='full' width='75%'  > 
                                        <Text>Partido</Text>
                                            {Paises == null || Paises.Paises == undefined || Paises.Paises.Paises.length === 0 ? (
                                                <Select placeholder='Selecciona Países' borderColor='black'>
                                                    <option value='opcion1'>Honduras-Argentina Ejemplo</option>
                                                    <option value='opcion2'>Alemania-Rusia Ejemplo</option>
                                                </Select>
                                            ):(
                                                <Select borderColor='black' onChange={CambioPais} placeholder='Selecciona Países'>
                                                    {Paises.Paises.Paises.map((Paisino) => (
                                                        <option value={Paisino}>{Paisino}</option>
                                                    ))}
                                                </Select>
                                            )}
                                    </Box>
                                </GridItem>
                                <GridItem colStart={8} colEnd={15} rowStart={1} rowEnd={1}>
                                    <Box  margin='auto' marginBottom='-10%' h='full' width='75%'  > 
                                    <Text>Fase</Text>
                                        <Select  borderColor='black' onChange={CambioFase} placeholder='Selecciona una opcion'>
                                            <option value='1'>Octavos</option>
                                            <option value='2'>Cuartos</option>
                                            <option value='3'>Semifinal</option>
                                            <option value='4'>Final</option>
                                        </Select>
                                    </Box>
                                </GridItem>
                                
                                <GridItem colStart={15} colEnd={16} rowStart={1} rowEnd={3}>

                                </GridItem>
                                <GridItem colStart={7} colEnd={9} rowStart={3} rowEnd={4}>
                                    {
                                        //<Button onClick={() => Actualizar_Grafo()} w='full'>Ver</Button>
                                    }
                                    
                                </GridItem>
                            </Grid>
                        </GridItem>
                        <GridItem rowStart={2} rowEnd={9} >
                            <Grid h='full' templateColumns='repeat(16)' templateRows='repeat(8)'>
                                <GridItem colStart={1} colEnd={11} rowStart={0} rowEnd={1} >

                                </GridItem>
                                <GridItem colStart={2} colEnd={10} rowStart={1} rowEnd={7} bg='yellow'>
                                    
                                    {Datos == null || Datos.Fase.Predics.length === 0 ? (
                                        <>
                                            <Heading textAlign='center' fontSize='2xl'>Predicciones por los fans</Heading>
                                            <br></br>
                                            <Text>No hay datos</Text>
                                        </>
                                    ):(
                                        <>
                                            <Heading fontSize='2xl'>Predicciones por los fans</Heading>
                                            <br></br>
                                            {Datos.Fase.Predics.map((Prediccion) => (
                                                <>
                                                    <Text>Resultado: {Prediccion.Punteo}</Text>
                                                    <Text>Votos: {Prediccion.Votos}</Text>
                                                    <br></br>
                                                </>
                                            ))}
                                            {Datos == undefined || Datos.Fase.Predics == null || Datos == null || data == null ? (
                                                <></>
                                            ):(
                                                <Box h='full' w='full'>
                                                    <Bar options={options} data={data} height={5} width={30}/>
                                                </Box>
                                                
                                            )}
                                            
                                        </>
                                    )}
                                </GridItem>
                            </Grid>
                        </GridItem>

                    </Grid>
                </GridItem>
            </Grid>
        </>
    );
}

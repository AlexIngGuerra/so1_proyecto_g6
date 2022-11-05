import { Box, Button, Grid, GridItem, Heading, Select, Table, TableContainer, Text } from '@chakra-ui/react';
import { React, useState, useEffect } from 'react';

import * as API from './Servicio/LogsServicio';



export default function Logs() {
  const [Logses, setLogses] = useState(null);
  const [Actualizar, setActualizar] = useState(true);
  const [algo, setalgo] = useState(null);
  useEffect(() => {
    API.getLogs().then(setLogses).catch(console.log);
  }, [Actualizar]);
  var alto = window.innerHeight;

  const Borrar = () => {
    API.deleteLogs().then(algo).catch(console.log);
    setActualizar(!Actualizar);
  }
  return (
    <>
      <Grid h={alto} templateRows='repeat(1)'
        templateColumns='repeat(9)' gap={1}>
        <GridItem colStart={1} colEnd={10} rowStart={1} rowEnd={1}>
          <Grid h='full' templateRows='repeat(9, fr)' templateColumns='repeat(17)' gap={1}>
            <GridItem rowStart={0} rowEnd={1} >
              <Grid h='full'>
                <Button bg='green.100' onClick={() => setActualizar(!Actualizar)}>
                  Actualizar
                </Button>
                <Button bg='blue' color='white' onClick={() => Borrar()}>
                  Limpiar Logs
                </Button>
                {Logses == null ? (
                  <></>
                ) : (
                  <Text>
                    CANTIDAD DE LOGS {Logses.Cantidad}
                  </Text>
                )}
                
              </Grid>
            </GridItem>
            <GridItem colStart={0} colEnd={1} rowStart={1} rowEnd={3}>
            
            </GridItem>
            <GridItem colStart={1} colEnd={8} h='full' rowStart={1} rowEnd={1} overflowY='auto'>
            
              {Logses == null || Logses.Logs == null || Logses.Logs.length == 0 ? (
                <Box rounded='md'>
                  No hay Logs
                </Box>
              ) : (
                <>
                  {Logses.Logs.map((Logito) => (
                    <Box rounded='md' border='1px' w='md'>
                      Team1: {Logito.team1}<br></br>
                      Team2: {Logito.team2}<br></br>
                      Score: {Logito.score}<br></br>
                      Fase: {Logito.phase}<br></br>
                    </Box>
                  ))}
                </>
              )}
            </GridItem>
          </Grid>
        </GridItem>
      </Grid>
    </>
  );
}

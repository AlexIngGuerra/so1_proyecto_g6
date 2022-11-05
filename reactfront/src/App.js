import logo from './logo.svg';
import './App.css';
import Live from './componentes/redis/Live';
import Logs from './componentes/mongo/Logs';
import { Routes, Route, Link, BrowserRouter, Navigate, useNavigate } from "react-router-dom";
import {React, useState, useEffect} from 'react';
import { Box, Button, Grid, GridItem, Heading, Select, Table, TableContainer, Text } from '@chakra-ui/react';

function App() {
  const navigate = useNavigate();
  var alto = window.innerHeight;

  const Vista_Live = () => {
    let path = `/live`
    navigate(path);
  }

  const Vista_Logs = () => {
    let path = `/logs`
    navigate(path);
  }

  return (
    <>
      <Grid h={alto} templateRows='repeat(1)'
        templateColumns='repeat(9)' gap={1}>
        <GridItem colStart={0} colEnd={1} rowStart={1} rowEnd={1} rounded='md' border='1px'>
          <Button onClick={Vista_Live}>Live</Button>
          <Button onClick={Vista_Logs}>Logs</Button>
        </GridItem>
        <GridItem colStart={1} colEnd={10} rowStart={1} rowEnd={1}>
          <Grid h='full' templateRows='repeat(9, fr)' templateColumns='repeat(17)' gap={1}>
            <GridItem rowStart={0} rowEnd={1} >
              <Grid h='full'>
                <Heading margin='auto' size='3xl'>
                  USACTAR
                </Heading>
              </Grid>
            </GridItem>
            <GridItem>
            <Routes>
              <Route path="/live" element={<Live />}></Route>
              <Route path="/logs" element={<Logs />}></Route>
            </Routes>
            </GridItem>
          </Grid>
        </GridItem>
      </Grid>
      
    </>
  );
}

export default App;

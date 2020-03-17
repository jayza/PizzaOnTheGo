import React, { useState, useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import { PizzaContext } from "./contexts/PizzaContext"
import { StateProps, Size } from './interfaces';
import Pizzas from './components/Pizzas';

const apiUrl = 'http://localhost:8080/api/v1';



const App = () => {
  const initialState: StateProps = {
    pizzas: [],
    sizes: [],
    crusts: [],
  }

  const [globalState, setGlobalState] = useState(initialState);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getInitialData()
  },[])

  const getInitialData = async () => {
    console.log("fetch pizzas")
    const pizzas = await fetch(apiUrl + '/pizzas?withIngredients=1')
      .then(res => res.json())
        .then(({data}) => data)
        .catch(console.error)
    
    
    console.log("fetch sizes")
    const sizes = await fetch(apiUrl + '/sizes')
    .then(res => res.json())
      .then(({ data }) => data)
      .catch(console.error)

    console.log("fetch crusts")
    const crusts = await fetch(apiUrl + '/crusts')
    .then(res => res.json())
      .then(({data}) => data)
      .catch(console.error)

    setGlobalState({
      ...globalState, 
      pizzas, 
      sizes, 
      crusts,
      defaultSize: sizes[0],
      defaultCrust: crusts[0]
    });

    console.log("set loading to false")
    setLoading(false)

  }

  return (
    <PizzaContext.Provider value={globalState}>
      {!loading &&(
      <>
        <h1>Pizzas</h1>
        <Pizzas/>
      </>
      )}
    </PizzaContext.Provider>
  );
}

export default App;

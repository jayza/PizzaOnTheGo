import React, { useState, useEffect } from 'react';
import './App.css';
import { PizzaContext } from "./contexts/PizzaContext"
import { StateProps, LineItem } from './interfaces';
import Pizzas from './components/Pizzas';
import {
  BrowserRouter as Router,
  Switch,
  Route
} from "react-router-dom";
import Cart from './components/Cart';
import Checkout from './components/Checkout';

const apiUrl = 'http://localhost:8080/api/v1';
const App = () => {
  const initialState: StateProps = {
    pizzas: [],
    sizes: [],
    crusts: [],
    toppings: [],
    lineItems: []
  }

  const [globalState, setGlobalState] = useState(initialState);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    getInitialData()
    // eslint-disable-next-line
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

    console.log("fetch crusts")
    const toppings = await fetch(apiUrl + '/toppings')
    .then(res => res.json())
      .then(({data}) => data)
      .catch(console.error)

    setGlobalState({
      ...globalState, 
      pizzas, 
      sizes, 
      crusts,
      toppings,
      defaultSize: sizes[0],
      defaultCrust: crusts[0]
    });

    console.log("set loading to false")
    setLoading(false)

  }

  return (
    <PizzaContext.Provider value={{
      globalState: globalState, 
      setLineItems: (lineItem: LineItem) => {
        setGlobalState({...globalState, lineItems: [...globalState.lineItems, lineItem]})
        console.log(globalState);
      }
    }}>
      {!loading &&(
      <Router>
        <Switch>
          <Route path="/checkout">
            <Checkout/>
          </Route>
          <Route path="/">
            <>
              <h1>Pizzas</h1>
              <Cart/>
              <Pizzas/>
            </>
          </Route>
        </Switch>
      </Router>
      )}
    </PizzaContext.Provider>
  );
}

export default App;

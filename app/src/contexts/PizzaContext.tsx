import React from 'react';
import { LineItem, StateProps } from '../interfaces'
// create context provider and consumer
export const PizzaContext = React.createContext<{globalState: StateProps, setLineItems: Function}>({
  globalState: {
    pizzas: [],
    sizes: [],
    crusts: [],
    toppings: [],
    lineItems: []
  },
  setLineItems: (lineItem: LineItem) => {}
});
import React from 'react';
import { StateProps } from '../interfaces'
// create context provider and consumer
export const PizzaContext = React.createContext<StateProps|null>(null);
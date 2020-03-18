import React from 'react'
import { PizzaContext } from '../contexts/PizzaContext';
import { Link } from 'react-router-dom';

const Cart = () => {
  return (
    <PizzaContext.Consumer>
      {(value) => (
        <>
          <div>{value.globalState.lineItems.length} items in the cart.</div>
          <Link to="/checkout">Go to checkout</Link>
        </>
      )}
    </PizzaContext.Consumer>
  )
};

export default Cart;
import React, { useState, useContext } from 'react'
import { PizzaContext } from '../contexts/PizzaContext';
import { LineItem, Order } from '../interfaces';
import { Link } from 'react-router-dom';
const apiUrl = 'http://localhost:8080/api/v1';
const Checkout = () => {
  const { globalState } = useContext(PizzaContext)

  let initialOrder: Order = {
    userId: 1, //temporary "logged in" user
    shippingInformation: {
      firstName: "",
      lastName: "",
      phone: "",
      streetAddress: "",
      zipCode: "",
      city: ""
    },
    lineItems: [...globalState.lineItems]
  }

  const [orderState, setOrderState] = useState(initialOrder);
  const [completeOrder, setCompleteOrder] = useState();

  const getUnitPrice = (lineItem: LineItem) => {
    let price = lineItem.item.price;

    if (lineItem.size !== undefined) {
      price += lineItem.size.price
    }

    if (lineItem.variation !== undefined) {
      price += lineItem.variation.price
    }

    if (lineItem.ingredients.length > 0) {
      lineItem.ingredients.forEach((i) => {
        price += i.price
      });
    }

    return price
  }

  const grandTotal = (lineItems: LineItem[]) => {
    let price = 0;

    lineItems.forEach((l) => {
      price += getUnitPrice(l) * l.quantity;
    })

    return price;
  }

  const handleChange = (e: any) => {
    setOrderState({
      ...orderState,
      shippingInformation: {
        ...orderState.shippingInformation,
        [e.target.name]: e.target.value
      }
    })
  }

  const handleSubmit = (e: any) => {
    e.preventDefault();

    postOrder().then(r => {
      setCompleteOrder({...r.data});
      setOrderState(initialOrder);
    }).catch(console.error)
    
  }

  const postOrder = async () => {
    console.log("posting", JSON.stringify(orderState))
    // Default options are marked with *
    const response = await fetch(apiUrl + '/orders?loggedInAs=1', {
      method: 'POST', // *GET, POST, PUT, DELETE, etc.
      mode: 'cors', // no-cors, *cors, same-origin
      // cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(orderState) // body data type must match "Content-Type" header
    });
    return await response.json(); // parses JSON response into native JavaScript objects
  }

  const receiptUrl = (order?: Order) => {
    console.log(order);
    if (order !== undefined) {
      return apiUrl + '/orders/' + order.id + '/receipt?loggedInAs=1';
    }
    else {
      return "";
    }
  }
  
  return (
    <PizzaContext.Consumer>
      {(value) => value && value.globalState !== undefined && value.globalState.lineItems !== null && (
        <>
          {(completeOrder === undefined) && (
          <div>
            <h1>Your order!</h1>
            {value.globalState.lineItems.length === 0 && (
              <div><p>There are no items in your cart.</p>
              <Link to="/">Go back to store..</Link>
              </div>
            )} 
            {value.globalState.lineItems.length > 0 && (
              <div>
              <table>
                <thead>
                  <tr>
                    <th>Product</th>
                    <th>Size</th>
                    <th>Quantity</th>
                    <th>Unit Price</th>
                    <th>Total</th>
                  </tr>
                </thead>
                <tbody>
                {value.globalState.lineItems.map((lineItem) => (
                  <tr>
                    <td>{lineItem.item.name}</td>
                    <td>{(lineItem.size !== undefined) ? lineItem.size.name : ""}</td>
                    <td>{lineItem.quantity}</td>
                    <td>{getUnitPrice(lineItem).toString()} SEK</td>
                    <td>{(getUnitPrice(lineItem) * lineItem.quantity).toString()} SEK</td>
                  </tr>
                ))}
                </tbody>
              </table>
            <div><p>Total: {grandTotal(value.globalState.lineItems)} SEK</p></div>
            <div>
              <h2>Shipping Information:</h2>
              <form onSubmit={handleSubmit}>
                <p>First name:</p>
                <input type="text" required onChange={handleChange} name="firstName" value={orderState.shippingInformation.firstName}/><br/>
                <p>Last name:</p>
                <input type="text" required onChange={handleChange} name="lastName" value={orderState.shippingInformation.lastName}/><br/>
                <p>Phone number:</p>
                <input type="text" required onChange={handleChange} name="phone" value={orderState.shippingInformation.phone}/><br/>
                <p>Street Address:</p>
                <input type="text" required onChange={handleChange} name="streetAddress" value={orderState.shippingInformation.streetAddress}/><br/>
                <p>Zip Code:</p>
                <input type="text" required onChange={handleChange} name="zipCode" value={orderState.shippingInformation.zipCode}/><br/>
                <p>City:</p>
                <input type="text" required onChange={handleChange} name="city" value={orderState.shippingInformation.city}/><br/>
                <input type="submit" value="Order"/>
              </form>
            </div>
            </div>
            )}
          </div>
          )}
          {(completeOrder !== undefined && (
            <div>
              <h1>Your order is completed!</h1>
              <p>Thank you for ordering with Pizza on the Go! Grab your receipt below.</p>
              <a href={receiptUrl(completeOrder)}>Download receipt</a> or 
              <Link to="/">Go back to the pizza shop</Link>
            </div>
          ))}
        </>
      )}
    </PizzaContext.Consumer>
  )
}

export default Checkout;
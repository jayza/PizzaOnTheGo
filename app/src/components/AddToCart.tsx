import React, { useContext, useState } from 'react'
import { Size, Variation, Ingredient, LineItem } from '../interfaces'
import { PizzaContext } from '../contexts/PizzaContext'

export const AddToCart = (value: any) => {
  const {globalState, setLineItems } = useContext(PizzaContext);
  const initialLineItem: LineItem = { 
    item: value.pizza, 
    size: globalState?.defaultSize,
    variation: globalState?.defaultCrust,
    specialInstruction: "",
    quantity: 1,
    ingredients: []
  };

  const [lineItemState, setLineItemState] = useState(initialLineItem);

  const handleChange = (e: any) => {
    const inputValue = e.target.value
    let value = null;
    switch (e.target.name) {
      case "size":
      value = globalState?.sizes.find(s => s.id === Number(inputValue))
      break;
      case "variation":
      value = globalState?.crusts.find(c => c.id === Number(inputValue))
      break;
      case "ingredients":
      value = [...lineItemState.ingredients, globalState?.toppings.find(t => t.id === Number(inputValue))]
      break;
      case "specialInstruction":
      value = inputValue
      break;
      case "quantity":
      value = parseInt(inputValue)
      break;
    }

    return setLineItemState({
      ...lineItemState,
      [e.target.name]: value,
    });
  }

  const handleSubmit = (e: any) => {
    e.preventDefault();
    setLineItems(lineItemState)
    setLineItemState({...initialLineItem});
  }

  const selectSize = (sizes: Size[], defaultSize?: Size) => {
    return sizes.map(size => (
      <div>
        <label>
          {size.name}
          <input 
            type="radio" 
            required 
            value={size.id} 
            onChange={handleChange}
            name="size" 
            checked={(lineItemState.size !== undefined) ? lineItemState.size.id === size.id : false}
            />
          {size.price} SEK
        </label>
      </div>
    ))
  }
     
  const selectCrust = (crusts: Variation[], defaultCrust?: Variation) => {
    return crusts.map(crust => (
      <div>
        <label>
          {crust.name}
          <input
            type="radio" 
            required 
            value={crust.id} 
            onChange={handleChange}
            name="variation"
            checked={(lineItemState.variation !== undefined) ? lineItemState.variation.id === crust.id : false} 
          />
          {crust.price} SEK
        </label>
      </div>
    ))
  }

  const selectIngredients = (ingredients: Ingredient[]) => {
    return ingredients.map(ingredient => (
      <div>
        <label>
          {ingredient.name}
          <input 
            type="checkbox" 
            value={ingredient.id} 
            onChange={handleChange} 
            checked={(lineItemState.ingredients.findIndex((i) => i.id === ingredient.id)) !== -1}
            name="ingredients"
          />
          +{ingredient.price} SEK
        </label>
      </div>
    ))
  }

  return (
    <PizzaContext.Consumer>
    {(value) => value.globalState !== null && value && (
    <form onSubmit={handleSubmit}>
      <div className="form__item">
        <b>Size:</b><br/>
        {selectSize(value.globalState.sizes, value.globalState.defaultSize)}
      </div>
      <div className="form__item">
        <b>Crust:</b><br/>
        {selectCrust(value.globalState.crusts, value.globalState.defaultCrust)}
      </div>
      <div className="form__item">
        <b>Ingredients:</b><br/>
        {selectIngredients(value.globalState.toppings)}
      </div>
      <div className="form__item">
        <b>Quantity:</b><br/>
        <input type="number" name="quantity" defaultValue="1" onChange={handleChange}/>
      </div>
      <div className="form__item">
        <b>Special instructions:</b><br/>
        <textarea name="specialInstruction" onChange={handleChange}></textarea>
      </div>
      <input value="Add to Cart" type="submit" />
    </form>
  )}
  </PizzaContext.Consumer>
  )
}
export interface Pizza {
  id: number;
  name: string;
  price: number;
  base: Ingredient;
  toppings: Ingredient[];
  dough: Ingredient;
}

export interface Size {
  id: number;
  price: number;
  size: string;
}

export interface Variation {
  id: number;
  price: number;
  name: string;
}

export interface Ingredient {
  id: number;
  price: number;
  name: string;
}

export interface StateProps {
  pizzas: Pizza[];
  sizes: Size[];
  crusts: Variation[];
  defaultSize?: Size;
  defaultCrust?: Variation;
}

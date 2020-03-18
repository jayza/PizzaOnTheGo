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
  name: string;
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

export interface LineItem {
  id?: number;
  price?: number;
  item: Pizza;
  quantity: number;
  size?: Size;
  variation?: Variation;
  specialInstruction: string;
  ingredients: Ingredient[];
}
export interface ShippingInformation {
  firstName: string;
  lastName: string;
  phone: string;
  streetAddress: string;
  zipCode: string;
  city: string;
}

export interface Order {
  id?: number;
  shippingInformation: ShippingInformation;
  userId: number;
  lineItems: LineItem[];
}

export interface GlobalState {
  globalState: StateProps;
  setLineItems: React.Dispatch<React.SetStateAction<StateProps>>;
}

export interface StateProps {
  pizzas: Pizza[];
  sizes: Size[];
  crusts: Variation[];
  toppings: Ingredient[];
  lineItems: LineItem[];
  defaultSize?: Size;
  defaultCrust?: Variation;
}

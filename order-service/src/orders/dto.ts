export class CreateOrderDto {
  readonly customerId: string;
  readonly products: { productId: string; quantity: number }[];
}

export class UpdateOrderDto {
  customerId?: string;
  products: { productId: string; quantity: number }[];
}

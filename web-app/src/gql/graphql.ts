/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export enum BardState {
  Created = 'CREATED',
  Finished = 'FINISHED',
  Running = 'RUNNING'
}

export type Board = {
  __typename?: 'Board';
  availableRoles?: Maybe<Array<Maybe<Scalars['String']>>>;
  createdAt: Scalars['String'];
  finished: Scalars['Boolean'];
  full: Scalars['Boolean'];
  id: Scalars['ID'];
  name: Scalars['String'];
  orders?: Maybe<Array<Maybe<Order>>>;
  players?: Maybe<Array<Maybe<Player>>>;
  state: BardState;
};

export type Mutation = {
  __typename?: 'Mutation';
  addPlayer?: Maybe<Player>;
  createBoard?: Maybe<Board>;
  createOrder?: Maybe<Order>;
  deliverOrder?: Maybe<Response>;
  updateWeeklyOrder?: Maybe<Response>;
};


export type MutationAddPlayerArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  role?: InputMaybe<Role>;
};


export type MutationCreateBoardArgs = {
  name?: InputMaybe<Scalars['String']>;
};


export type MutationCreateOrderArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  receiverId?: InputMaybe<Scalars['String']>;
};


export type MutationDeliverOrderArgs = {
  amount?: InputMaybe<Scalars['Int']>;
  boardId?: InputMaybe<Scalars['String']>;
  orderId?: InputMaybe<Scalars['String']>;
};


export type MutationUpdateWeeklyOrderArgs = {
  amount?: InputMaybe<Scalars['Int']>;
  boardId?: InputMaybe<Scalars['String']>;
  playerId?: InputMaybe<Scalars['String']>;
};

export type Order = {
  __typename?: 'Order';
  amount: Scalars['Int'];
  board?: Maybe<Board>;
  createdAt?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  originalAmount: Scalars['Int'];
  receiver?: Maybe<Player>;
  sender: Player;
  state: OrderState;
  type: OrderType;
};

export enum OrderState {
  Delivered = 'DELIVERED',
  Pending = 'PENDING'
}

export enum OrderType {
  CpuOrder = 'CPU_ORDER',
  PlayerOrder = 'PLAYER_ORDER'
}

export type Player = {
  __typename?: 'Player';
  backlog: Scalars['Int'];
  board?: Maybe<Board>;
  cpu: Scalars['Boolean'];
  id: Scalars['ID'];
  lastOrder: Scalars['Int'];
  name: Scalars['String'];
  orders?: Maybe<Array<Order>>;
  role: Role;
  stock: Scalars['Int'];
  weeklyOrder: Scalars['Int'];
};

export type Query = {
  __typename?: 'Query';
  getBoard?: Maybe<Board>;
  getBoardByName?: Maybe<Board>;
  getPlayer?: Maybe<Player>;
  getPlayersByBoard?: Maybe<Array<Maybe<Player>>>;
};


export type QueryGetBoardArgs = {
  id?: InputMaybe<Scalars['String']>;
};


export type QueryGetBoardByNameArgs = {
  name?: InputMaybe<Scalars['String']>;
};


export type QueryGetPlayerArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  playerId?: InputMaybe<Scalars['String']>;
};


export type QueryGetPlayersByBoardArgs = {
  boardId?: InputMaybe<Scalars['String']>;
};

export type Response = {
  __typename?: 'Response';
  message?: Maybe<Scalars['String']>;
  status?: Maybe<Scalars['Int']>;
};

export enum Role {
  Factory = 'FACTORY',
  Retailer = 'RETAILER',
  Wholesaler = 'WHOLESALER'
}

export type Subscription = {
  __typename?: 'Subscription';
  board?: Maybe<Board>;
  newOrder?: Maybe<Order>;
  orderDelivery?: Maybe<Order>;
  player?: Maybe<Player>;
};


export type SubscriptionBoardArgs = {
  boardId?: InputMaybe<Scalars['String']>;
};


export type SubscriptionNewOrderArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  playerId?: InputMaybe<Scalars['String']>;
};


export type SubscriptionOrderDeliveryArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  playerId?: InputMaybe<Scalars['String']>;
};


export type SubscriptionPlayerArgs = {
  boardId?: InputMaybe<Scalars['String']>;
  playerId?: InputMaybe<Scalars['String']>;
};

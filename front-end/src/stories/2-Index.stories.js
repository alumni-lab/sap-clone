import React, { Fragment } from 'react'

import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";


import InventoryListItem from "../components/InventoryListItem";
import InventoryList from "../components/InventoryList";
import Header from "../components/Header";
import PurchaseOrderItem from "../components/PurchaseOrderItem";
import PurchaseOrderList from "../components/PurchaseOrderList";

const items = [
  {
    id: 1,
    name: "Ribs",
    quantity: 2,
    description: "saucy goodness",
    type: "type",
    vendor: "Slow Cow Farms"
  },
  {
    id: 2,
    name: "paper planes",
    quantity: 5,
    description: "not for dipping in sauce",
    type: "type",
    vendor: "Dunder Mifflin"
  },
  {
    id: 3,
    name: "frogs",
    quantity: 11,
    description: "Rrrbit",
    type: "type",
    vendor: "a pond near here"
  },
];

const POs = [
  {
    id: 1,
    POid: 1000,
    customerName: "BigBoxGuy",
    amount: 12000
  },
  {
    id: 2,
    POid: 1005,
    customerName: "SmallBoxGuy",
    amount: 125000
  },
  {
    id: 3,
    POid: 1050,
    customerName: "Mr. Andersen",
    amount: 10110
  } 
];

storiesOf("InventoryListItem", module) //Initiates Storybook and registers our InventoryListItem component
.addParameters({
  backgrounds: [{ name: "dark", value: "#222f3e", default: true }]
}) // Provides the default background color for our component
.add("Unselected", () => <InventoryListItem name="Rack of Ribs" quantity={5} description="saucy happiness" type="type" vendor="Slow Cow Farms"/>)
.add("Selected", () => <InventoryListItem name="Rack of Ribs" quantity={5} description="saucy happiness" type="type" vendor="Slow Cow Farms"selected />) 


storiesOf("InventoryList", module)
  .add("InventoryList", () => (
    <InventoryList items={items} />
  ))

storiesOf("Header", module)
.add("logged in", () => (
  <Header username="Bobby" logo="https://i.imgur.com/LpaY82x.png"/>
))
.add("logged out", () => (<Header logo="https://i.imgur.com/LpaY82x.png"/>))

storiesOf("PurchaseOrderItem", module) //Initiates Storybook and registers our InventoryListItem component
.addParameters({
  backgrounds: [{ name: "dark", value: "#222f3e", default: true }]
}) // Provides the default background color for our component
.add("Unselected", () => <PurchaseOrderItem customerName="BigBoxVender" POid={1} amount={1000.00} />)

storiesOf("PurchaseOrderList", module)
.add("PurchaseOrderList", () => (
  <PurchaseOrderList POs={POs} />
))


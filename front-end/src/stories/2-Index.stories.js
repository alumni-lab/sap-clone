import React, { Fragment } from 'react'

import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";


import InventoryListItem from "../components/InventoryListItem";
import InventoryList from "../components/InventoryList";


storiesOf("InventoryListItem", module) //Initiates Storybook and registers our InventoryListItem component
.addParameters({
  backgrounds: [{ name: "dark", value: "#222f3e", default: true }]
}) // Provides the default background color for our component
.add("Unselected", () => <InventoryListItem name="Rack of Ribs" quantity={5} description="saucy happiness" type="type" vendor="Slow Cow Farms"/>) // To define our stories, we call add() once for each of our test states to generate a story
.add("Selected", () => <InventoryListItem name="Rack of Ribs" quantity={5} description="saucy happiness" type="type" vendor="Slow Cow Farms"selected />) 

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

storiesOf("InventoryList", module)
  .add("InventoryList", () => (
    <InventoryList items={items} />
  ))
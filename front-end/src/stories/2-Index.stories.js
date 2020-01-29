import React, { Fragment } from 'react'

import { storiesOf } from "@storybook/react";
import { action } from "@storybook/addon-actions";


import InventoryListItem from "../components/InventoryListItem";


storiesOf("InventoryListItem", module) //Initiates Storybook and registers our InventoryListItem component
.addParameters({
  backgrounds: [{ name: "dark", value: "#222f3e", default: true }]
}) // Provides the default background color for our component
.add("Unselected", () => <InventoryListItem name="Rack of Ribs" items={5} />) // To define our stories, we call add() once for each of our test states to generate a story
.add("Selected", () => <InventoryListItem name="Rack of Ribs" items={5} selected />) 


import React from 'react';
// import InventoryList from "./components/InventoryList";
import Dashboard from "./components/Dashboard";
import Header from "./components/Header";
import Footer from "./components/Footer";

function App() {
  //Temp data
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

  return (
    <h1>
      <Header />
      <Dashboard items={items} />
      <Footer />
    </h1>
  );
}

export default App;

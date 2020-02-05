# Components

## App component to render the page
  1. Header component contains username, loggedIn?, company logo
  2. Dashoard component contains page title and has 2 views.
  * button to add Customer renders Add Customer component (moodle)
  *  View 1 Inventory
      * Add item button renders AddItem component (moodle)
      * searchbar 
      * table column titles (item id, description, qty, etc)
      * InventoryList component to display InventoryListItems based on which table column is selected
        * InventoryListItem (each InventoryListItem renders with an edit button)
  
  * View 2 Purchase Orders
    * Create purchase order button renders purchase order component (moodle)
    * searchbar
    * table column titles (PO id, customer name, amount)
    * PurchaseOrderList component to display PurchaseOrders based on which table column is selected
      *  PurchaseOrder (renders with an edit button)
  3. Footer component with previous, next, page numbers

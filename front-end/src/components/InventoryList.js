// import React from "react";
// import ReactDOM from "react-dom";
// import InventoryListItem from "../components/InventoryListItem.js";

// export default function InventoryList(props) {


//   const items = props.items.map(item => {
//     return (
//       <InventoryListItem
//         key={item.id}
//         name={item.name}
//         quantity={item.quantity}
//         description={item.description}
//         price={item.price}
//         type={item.type}
//         vendor={item.vendor}
//       />
//     );
//   });

//   return items;
// }

// import React from "react";
// import "../components/InventoryListItem.scss";


// export default function InventoryListItem(props) {
//   return (
//     <li>
//       <h2 >{props.name}</h2>
//       <h3 >{props.quantity}</h3> 
//       <p >{props.description}</p>
//       <p >{props.type}</p>
//       <p >{props.price}</p>
//       <h3 >{props.vendor}</h3>
//     </li>
//   );
// }

import React from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';

const useStyles = makeStyles({
  table: {
    minWidth: 650,
  },
});


export default function SimpleTable(props) {
  const classes = useStyles();

  return (
    <TableContainer component={Paper}>
      <Table className={classes.table} aria-label="simple table">
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell align="right">description</TableCell>
            <TableCell align="right">quantity</TableCell>
            <TableCell align="right">price</TableCell>
            <TableCell align="right">type</TableCell>
            <TableCell align="right">vendor</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {props.items.map(item => (
            <TableRow key={item.name}>
              <TableCell component="th" scope="row">
                {item.name}
              </TableCell>
              <TableCell align="right">{item.description}</TableCell>
              <TableCell align="right">{item.quantity}</TableCell>
              <TableCell align="right">{item.price}</TableCell>
              <TableCell align="right">{item.type}</TableCell>
              <TableCell align="right">{item.vendor}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
}
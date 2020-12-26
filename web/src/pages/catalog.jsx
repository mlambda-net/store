

import React from "react";
import {connect} from "react-redux";
import {withStyles, withTheme} from "@material-ui/core/styles";
import {Action, State} from "@mlambda-net/core/packages/common";
import ProductList from "./catalog/product_list";
import CreateProduct from "./catalog/create_product";


const styles = (theme) => ({

})

class Catalog extends React.Component {

  constructor(props) {
    super(props);
    this.state = {actual: "search"}
  }

  render() {
    return (
      <Action actual={this.state.actual}>
        <State name="search">
          <ProductList onAdd={() => this.setState({actual: "add"})}
                       onEdit={() => this.setState({actual: "edit"})}
                       onDelete={()=>{}}/>
        </State>
        <State name="add">
          <CreateProduct onClose={()=> this.setState({actual:"search"})}/>
        </State>
        <State name="edit">
          <div>Edit</div>
        </State>
      </Action>
    )
  }
}

export default connect(s => s.catalog)(withStyles(styles)(withTheme(Catalog)));

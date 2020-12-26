import React from "react";
import { withTheme} from "@material-ui/core/styles";
import {Divider, IconButton, InputAdornment, Paper, TextField, Typography} from "@material-ui/core";
import { DataGrid} from '@material-ui/data-grid';
import LocalizedStrings from "react-localization";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import {Add, Delete, Edit, Search} from "@material-ui/icons";
import CatalogService from "../../services/catalog";
import {connect} from "react-redux";

const lang = new LocalizedStrings({
  en: {
    title: 'Catalog Search',
    search: 'Search',
    code: 'Code',
    name: 'Name',
    description: 'Description',
    price: 'Price'
  },
  es: {
    title: 'Busqueda de productos',
    search: 'Buscar',
    code: 'Código',
    name: 'Nombre',
    description: 'Descripción',
    price: 'Precio'
  },
});


const styles = (theme) => ({
    view: {
      height: '100%',
      borderRadius: '10px'
    },

    viewTitle: {
      height: '46px',
      padding: theme.spacing(2, 2),
      alignItems: 'center',
      color: theme.palette.secondary.main,
      display: 'flex',
    },

    viewContent: {
      margin: theme.spacing(2, 2),
      height: 'calc(100% - 132px);'
    },

    viewActions: {
      height: '46px',
      display: 'flex',
      alignItems:'center',
      padding: theme.spacing(2,0)
    },

    viewActionSpace: {
      flexGrow:'2'
    },

    viewActionSearch: {
      flexGrow: '1'
    },

    viewGrid : {
      height: 'calc(100% - 56px)'
    }
})

class ProductList extends React.Component {

  constructor(props) {
    super(props);
    this.dispatch = props.dispatch
    this.service = new CatalogService(this.dispatch)
  }

  async componentDidMount() {
    this.service.fetch("")
  }

  render() {
    const columns = [
      {field: "code", headerName: lang.code, type: 'string', width: 200},
      {field: "name", headerName: lang.name, type: 'string', width: 300},
      {field: "description", headerName: lang.description, type: 'string', width: 500},
      {field: "priceText", headerName: lang.price, type: 'number', width: 150},
    ]
    const {classes, items} = this.props
    return (
      <Paper className={classes.view} elevation={3}>
        <div className={classes.viewTitle}>
          <Typography variant="h6">
            {lang.title}
          </Typography>
        </div>
        <Divider variant="fullWidth"/>
        <div className={classes.viewContent}>
          <div className={classes.viewActions}>
            <IconButton color="secondary" onClick={this.props.onAdd}>
              <Add/>
            </IconButton>
            <IconButton color="secondary" onClick={this.props.onEdit}>
              <Edit/>
            </IconButton>
            <IconButton color="secondary" onClick={this.props.onDelete}>
              <Delete/>
            </IconButton>
            <div className={classes.viewActionSpace}/>
            <TextField className={classes.viewActionSearch} label={lang.search}
                       color="secondary" size="small"
                       InputProps={{
                         endAdornment: <InputAdornment position="end"><Search color="secondary"/></InputAdornment>,
                       }}/>
          </div>
          <div className={classes.viewGrid}>
            <DataGrid columns={columns} rows={items} pageSize={10} style={{height: '100%'}}/>
          </div>
        </div>
      </Paper>
    )
  }
}


export default connect(s => s.catalog)(withUtils(styles)(withTheme(ProductList)));

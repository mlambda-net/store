import React from "react";
import Box from "@material-ui/core/Box";
import {
  Button, Divider,
  Grid, IconButton,
  InputAdornment,
  Paper,
  TextField,
  Typography
} from "@material-ui/core";
import SaveIcon from "@material-ui/icons/Save";
import { withTheme} from "@material-ui/core/styles";
import { v4 as uuidv4 } from 'uuid';
import {connect} from "react-redux";
import CatalogService from "../../services/catalog";
import {Cancel} from "@material-ui/icons";
import CurrencySelect from "../../components/currency"
import Brand from "../../components/brand"
import Category from "../../components/category"
import Preview from "../../components/preview"
import Images from "../../components/images"
import LocalizedStrings from "react-localization";
import withUtils from "@mlambda-net/core/packages/utils/withUtils";
import Template from "../../assets/template.jpg"

const lang = new LocalizedStrings({
  en: {
    title: "Add Product",
    code: "Code",
    name: "Name",
    price: "Price",
    description: "Description",
  },
  es: {
    title: "Agrega el producto",
    code: "Código",
    name: "Nombre",
    price: "Precio",
    description: "Descripción",
  },
});

const styles = (theme) => ({
  view: {
    minHeight: '100%',
    display: 'flex',
    flexDirection: 'column',
    borderRadius: '10px'
  },

  viewTitle: {
    height: '46px',
    padding: theme.spacing(2, 2),
    color: theme.palette.secondary.main,
    display: 'flex',
    alignItems: 'center'
  },

  title : {
    flexGrow: "1"
  },

  viewContent: {
    margin: theme.spacing(2, 2),
    flexGrow: "1"
  },

  viewActions: {
    height: '56px',
    display: 'flex',
    alignItems:'center',
    justifyContent: 'left'
  }
})

class CreateProduct extends React.Component {

  constructor(props) {
    super(props);
    this.state = {
      id: uuidv4(),
      open: false,
      currency: "USD",
      category: "",
      brand: "",
      preview: Template,
      images: [Template, Template, Template, Template]
    }
    this.dispatch = props.dispatch;
    this.service = new CatalogService(this.dispatch)
    this.selectedCurrency = this.selectedCurrency.bind(this)
    this.categoryChange = this.categoryChange.bind(this)
    this.brandChange = this.brandChange.bind(this)
    this.previewChange = this.previewChange.bind(this)
    this.imagesChange = this.imagesChange.bind(this)
    this.saveProduct = this.saveProduct.bind(this)
  }

  selectedCurrency(e) {
    this.setState({currency: e.target.value})
  }

  brandChange(v) {
    this.setState({brand: v})
  }

  categoryChange(e) {
    this.setState({category: e.target.value})
  }

  previewChange(img) {
    this.setState({preview: img})
  }

  imagesChange(img) {
    this.setState({images: img})
  }

  saveProduct() {
    this.service.save({
      id: this.state.id,
      code: this.state.code,
      name: this.state.name,
      currency: this.state.currency,
      price: this.state.price,
      category: this.state.category,
      brand: this.state.brand,
      description: this.state.description,
      thumbnail: this.state.preview,
      images: this.state.images
    })
  }

  render() {
    const {classes, onClose} = this.props
    lang.setLanguage(this.props.lang)

    return (
      <Paper elevation={3} className={classes.view}>
        <Box padding={2} className={classes.viewTitle}>
          <Typography variant="h6" className={classes.title}>
            {lang.title}
          </Typography>
          <IconButton color="secondary" onClick={onClose}>
            <Cancel/>
          </IconButton>
        </Box>
        <Divider/>
        <Box padding={2} className={classes.viewContent}>
          <Grid container spacing={3}>

            <Grid item xs={6} sm={4}>
              <TextField id="code" variant="standard" color="secondary" label={lang.code} fullWidth
                         value={this.state.code} onChange={e => this.setState({code: e.target.value})}/>
            </Grid>

            <Grid item xs={6} sm={4}>
              <TextField id="name" variant="standard" color="secondary" label={lang.name} fullWidth
                         value={this.state.name} onChange={e => this.setState({name: e.target.value})}/>
            </Grid>

            <Grid item xs={6} sm={4}>
              <Brand value={this.state.brand} onChange={this.brandChange}/>
            </Grid>

            <Grid item xs={6} sm={4}>
              <CurrencySelect value={this.state.currency} onCurrency={(this.selectedCurrency)}/>
            </Grid>

            <Grid item xs={6} sm={4}>
              <TextField id="price" label={lang.price} variant="standard" color="secondary"
                         type="number" fullWidth onChange={e => this.setState({price: e.target.value})}
                         value={this.state.price}
                         InputProps={{
                           startAdornment: <InputAdornment
                             position="start">{this.state.currency === "USD" ? "$" : "₡"}</InputAdornment>
                         }}
              />
            </Grid>

            <Grid item xs={6} sm={4}>
              <Category value={this.state.category} onChange={this.categoryChange}/>
            </Grid>

            <Grid item xs={12}>
              <TextField id="description" label={lang.description} color="secondary" variant="standard"
                         multiline rows={4} fullWidth
                         value={this.state.description} onChange={e => this.setState({description: e.target.value})}
              />
            </Grid>

            <Grid item xs={12} sm={12} md={5} lg={3}>
              <Preview id={this.state.id} value={this.state.preview} onChange={this.previewChange}/>
            </Grid>

            <Grid item xs={12} sm={12} md={7} lg={9}>
              <Images id={this.state.id} value={this.state.images} onChange={this.imagesChange}/>
            </Grid>

          </Grid>
        </Box>
        <Divider/>
        <Box padding={2} className={classes.viewActions}>
          <Button variant="contained" color="secondary" size="small" className={classes.button} startIcon={<SaveIcon/>} onClick={this.saveProduct}>
            Save
          </Button>
        </Box>
      </Paper>
    )
  }
}

export default  connect(s => s.catalog)(withUtils(styles)(withTheme(CreateProduct)));

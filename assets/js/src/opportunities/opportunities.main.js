import React, {Component} from 'react'
import List from '@material-ui/core/List'
import ListItem from "@material-ui/core/ListItem";
import axios from 'axios'
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import AddIcon from "@material-ui/icons/Add"
import SaveIcon from "@material-ui/icons/Save"
import EditIcon from "@material-ui/icons/Edit"
import OpportunityDetail from "./oportunities.detail";

const hostUri = "http://localhost:3000/api";

const OpportunityList = props => {
    const {select, opportunities} = props

    return (
        <List>
            {opportunities.map(item => {
                return (
                    <ListItem button onClick={_ => select(item)} key={item.id}>
                        {item.name}
                        |
                        {item.summary}
                    </ListItem>
                )
            })}
        </List>
    )
}

// const OpportunityDetail = props => {
//     const {detail} = props;
//     return (
//         <List>
//             {Object.keys(detail)
//                 .filter(x => _.includes(['name', 'summary', 'business_category'], x))
//                 .map(key => {
//                     return (
//                         <ListItem key={key}>
//                             <TextField
//                                 id={key}
//                                 label={key}
//                                 value={detail[key]}
//                                 onChange={data => props.handleChange(key, data)}
//                                 margin="normal"
//                                 disabled={!props.editing}
//                             />
//                         </ListItem>
//                     )
//                 })}
//         </List>
//     )
// }


class Opportunities extends Component {

    constructor(props) {
        super(props);
        this.state = {
            opportunities: [{name: 'hello world', id: 1}, {name: "wornderful", id: 2}],
            detail: {
                metrics: [],
                fatal_attributes: []
            },
            editing: false,
            updating: false,
            metrics: [],
            fatalAttributes: []
        }
    }

    componentDidMount() {


        this._getOpportunities();
        this._getMetrics();
        this._getFatalAttributes();
        this._getValues()

    }

    _getOpportunities() {
        axios({
            method: "get",
            url: `${hostUri}/opportunities`,
            responseType: "application/json",
        })
            .then(data => {

                this.setState({
                    opportunities: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })
    }

    _getMetrics = () => {
        axios({
            method: "get",
            url: `${hostUri}/metrics`,
            responseType: "application/json",
        })
            .then(data => {
                this.setState({
                    metrics: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })
    }
    _getValues = () => {
        axios({
            method: "get",
            url: `${hostUri}/values`,
            responseType: "application/json",
        })
            .then(data => {
                this.setState({
                    values: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })
    }


    _getFatalAttributes = () => {
        axios({
            method: "get",
            url: `${hostUri}/fatal_attributes`,
            responseType: "application/json",
        })
            .then(data => {
                this.setState({
                    fatalAttributes: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })
    }


    _newOpp = () => {

        // TODO make sure all new opportunities have all the MEtrics added to them available
        this.setState({
            editing: true,
            updating: false
        })
        axios({
            method: "get",
            url: `${hostUri}/opportunities/new`,
            responseType: "application/json",
        })
            .then(data => {
                let d = {...data.data}
                d.metrics = JSON.parse(
                    JSON.stringify(
                        this.state.metrics))
                    .map(item => {
                        delete item.id
                        return item
                    })
                this.setState({


                    detail: d
                })
            })
            .catch(error => {
                console.log(error)
            })
    }

    _saveDetail = () => {
// better handling is needed


        //TODO Handle the creation of the saving of metrics via the opportunity metric endpoints.


        if (!this.state.editing) {
            this.setState({
                editing: true
            });
            return
        }

        if (this.state.updating) {
            axios({
                method: "put",
                url: `${hostUri}/opportunities/${this.state.detail.id}`,
                responseType: "application/json",
                data: this.state.detail
            })
                .then(data => {
                    this.setState({
                        editing: false
                    })
                })
                .catch(error => {
                    console.log(error)
                })
        } else {

            axios({
                method: "post",
                url: `${hostUri}/opportunities`,
                responseType: "application/json",
                data: this.state.detail
            })
                .then(data => {
                    this.setState({
                        editing: false
                    })
                })
                .catch(error => {
                    console.log(error)
                })

        }
    }

    _handleChange = (key, data, index) => {
            let value = data.target.value;
        if (key === "value") {

            this.state.detail.metrics[index].value = value
            this.forceUpdate()

            return
        }
        this.setState((prevState, props) => {


            switch (key) {
                case "metric":
                    value = [...prevState.detail.metrics, value];
                    break;
                case "fatal_attribute":
                    debugger
                    value = [...(prevState.detail.fatal_attributes || []), value];
                    break;
                default:
            }


            return {
                detail: {
                    ...this.state.detail,
                    ... {[key]: value}
                }
            }
        })
        console.log(data)
    }

    _selectDetail = item => {
        this.setState({
            detail: item,
            updating: true
        })
    }

    _cancelEdit = _ => {
        this.setState({
            editing: false
        })
    }

    render() {
        const {fatalAttributes, metrics, values, editing, detail} = this.state
        return (<div>
                <Grid container spacing={24}>
                    <Grid item xs={6}>
                        <OpportunityList opportunities={this.state.opportunities} select={this._selectDetail}/>
                        <Button variant="fab" color="primary" aria-label="add" onClick={this._newOpp}>
                            <AddIcon/>
                        </Button>
                    </Grid>
                    <Grid item xs={6}>
                        <OpportunityDetail
                            detail={detail}
                            handleChange={this._handleChange}
                            editing={editing}
                            metrics={metrics}
                            values={values}
                            fatal_attributes={fatalAttributes}
                        />
                        <Button variant="fab" color="default" aria-label="add" onClick={this._saveDetail}>
                            {this.state.editing ? <SaveIcon/> : <EditIcon/>}
                        </Button>

                        {this.state.editing && (
                            <Button color="secondary" onClick={this._cancelEdit}>
                                cancel
                            </Button>
                        )}

                    </Grid>
                </Grid>
            </div>
        )
    }
}

const styles = theme => ({
    container: {
        display: 'flex',
        flexWrap: 'wrap',
    },
    textField: {
        marginLeft: theme.spacing.unit,
        marginRight: theme.spacing.unit,
        width: 200,
    },
    menu: {
        width: 200,
    },
});


export default Opportunities
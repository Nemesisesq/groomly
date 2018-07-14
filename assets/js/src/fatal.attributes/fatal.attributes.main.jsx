import React, {Component} from 'react'
import List from '@material-ui/core/List'
import ListItem from "@material-ui/core/ListItem";
import axios from 'axios'
import Grid from "@material-ui/core/Grid";
import Button from "@material-ui/core/Button";
import AddIcon from "@material-ui/icons/Add"
import SaveIcon from "@material-ui/icons/Save"
import EditIcon from "@material-ui/icons/Edit"
import TextField from "@material-ui/core/TextField";
import _ from 'lodash'

const hostUri = "http://localhost:3000/api";

const FatalAttributeList = props => {
    const {select, fatal_attributes} = props

    return (
        <List>
            {fatal_attributes.map(item => {
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

const FatalAttributeDetail = props => {
    const {detail} = props;
    return (
        <List>
            {Object.keys(detail)
                .filter(x => !_.includes(['id', 'created_at', 'updated_at'], x))
                .map(key => {
                    return (
                        <ListItem key={key}>
                            <TextField
                                id={key}
                                label={key}
                                value={detail[key]}
                                onChange={data => props.handleChange(key, data)}
                                margin="normal"
                                disabled={!props.editing}
                            />
                        </ListItem>
                    )
                })}
        </List>
    )
}


class FatalAttributes extends Component {

    constructor(props) {
        super(props);
        this.state = {
            fatal_attributes: [{name: 'hello world', id: 1}, {name: "wornderful", id: 2}],
            detail: {},
            editing: false,
            updating: false
        }
    }

    componentDidMount() {


        axios({
            method: "get",
            url: `${hostUri}/fatal_attributes`,
            responseType: "application/json",
        })
            .then(data => {

                this.setState({
                    fatal_attributes: [...data.data]
                })
            })
            .catch(error => {
                console.log(error)
            })

    }

    _newOpp = () => {
        this.setState({
            editing: true,
            updating: false
        })
        axios({
            method: "get",
            url: `${hostUri}/fatal_attributes/new`,
            responseType: "application/json",
        })
            .then(data => {
                this.setState({
                    detail: {...data.data}
                })
            })
            .catch(error => {
                console.log(error)
            })
    }

    _saveDetail = () => {
// better handling is needed
        if (!this.state.editing) {
            this.setState({
                editing: true
            });
            return
        }

        if (this.state.updating) {
            axios({
                method: "put",
                url: `${hostUri}/fatal_attributes/${this.state.detail.id}`,
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
                url: `${hostUri}/fatal_attributes`,
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
    _handleChange = (key, data) => {

        let real_data = data.target.value
        switch (key) {
            case 'weight':
                real_data = parseInt(real_data) || 0;
                break
        }

        this.setState({
            detail: {
                ...this.state.detail,
                ... {[key]: real_data}
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
        return (<div>
                <Grid container spacing={24}>
                    <Grid item xs={6}>
                        <FatalAttributeList fatal_attributes={this.state.fatal_attributes} select={this._selectDetail}/>
                        <Button variant="fab" color="primary" aria-label="add" onClick={this._newOpp}>
                            <AddIcon/>
                        </Button>
                    </Grid>
                    <Grid item xs={6}>
                        <FatalAttributeDetail detail={this.state.detail} handleChange={this._handleChange}
                                           editing={this.state.editing}/>
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


export default FatalAttributes
import React, {Component} from 'react'
import List from '@material-ui/core/List'
import ListItem from '@material-ui/core/ListItem'
import FormControl from "@material-ui/core/FormControl";
import InputLabel from "@material-ui/core/InputLabel";
import Select from "@material-ui/core/Select";
import MenuItem from "@material-ui/core/MenuItem";
import FormHelperText from "@material-ui/core/FormHelperText";
import Input from "@material-ui/core/Input";

export default class FatalAttributes extends Component {
    state = {

    }
    render() {
        const {detail, handleChange, editing, fatal_attributes} = this.props;
        const fa = detail.fatal_attributes || [];
        return (
            <List>
                {fa.map((item, index) => {
                    return (
                        <ListItem key={index}>
                            {` ${item.name} | ${item.summary}`}
                        </ListItem>
                    )
                })}


                <FormControl>
                    <InputLabel htmlFor="age-helper">this.stateAdd a Fatal Attribute</InputLabel>
                    <Select
                        value={this.state.fatalAttribute.name}
                        onChange={data => handleChange("fatal_attribute", data)}
                        input={<Input name="fatal_attribute" id="fatal_attribute-helper"/>}
                    >
                        {fatal_attributes.map(item2 => {
                            return (
                                <MenuItem value={item2}>{item2.name}</MenuItem>
                            )
                        })}
                    </Select>
                    <FormHelperText>Select a Fatal Attribute to Add</FormHelperText>
                </FormControl>
            </List>
        )
    }
}
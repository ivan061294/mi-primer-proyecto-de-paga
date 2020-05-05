import React, { Component } from 'react';

class SelectStatus extends Component {
    constructor(props) {
        super(props);
    }
    componentDidMount() {}
    render() {
        const {
            name, optionState, ...props
        } = this.props;
        const options = [
            {value:'A', label:'Aprobado'},
            {value:'P', label:'Pendiente'},
            {value:'R', label:'Rechazado'}
        ];
        let selected;
        if (optionState) {
            selected = options.find((o)=> (o.label===optionState)).value;
        }
        return (<select className="form-control" name={name} value={selected}>
                {options.map((o, idx)=>(<option key={idx} value={o.value}>{o.label}</option>))}
            </select>)
    }
}

export default SelectStatus;
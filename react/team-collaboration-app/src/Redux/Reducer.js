const initialState = 
{
    Projects : []
}

const fetchDataReducer = (state = initialState, action) => 
{
    switch(action.type)
    {
        case "fetchData" : 
        {

        }

        default : return state;
    }
}

export default fetchDataReducer;
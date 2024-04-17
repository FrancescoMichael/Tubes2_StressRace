function Button(props){
	const { text, onClick = () => {},} = props;
	return(
		<button 
			className= {`h-12 px-8 font-semibold rounded-3xl bg-blue-600 text-white`}
			type = "button"
			onClick = {onClick}
		> 
			{text}
			
        </button>
	);
}

export default Button
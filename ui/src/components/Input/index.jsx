import Prop from 'prop-types';

// eslint-disable-next-line react/prop-types
export const Input = ({ onChange, value }) => {
  return <input value={value || ''} onChange={onChange} type="text" />;
};

Input.propTypes = {
  onChange: Prop.func.isRequired,
};

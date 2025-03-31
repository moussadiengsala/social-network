const CustomSelect = ({
  label,
  placeholder,
  icon,
  options,
  name,
  onChange,
  register,
  error,
  additionalAttributes,
}) => {
  return (
    <div {...additionalAttributes}>
      <label htmlFor={name} className="sr-only">
        {label}
      </label>
      <div className="relative">
        <select
          {...register(name)}
          id={name}
          name={name}
          {...additionalAttributes}
          className="w-full rounded-lg border-gray-100 bg-white  border-2 p-4 pe-12 text-sm shadow-sm"
          placeholder={placeholder}
          onChange={onChange}
        >
          {options.map((opt, index) => (
            <option key={index} defaultValue={index == 0} value={opt}>
              {opt}
            </option>
          ))}
        </select>
        {icon && (
          <span className="absolute inset-y-0 end-0 grid place-content-center px-4">
            {icon}
          </span>
        )}
      </div>
      {error && (
        <div className="text-pink-500 mt-2 text-xs">*{error.message}</div>
      )}
    </div>
  );
};

export default CustomSelect;

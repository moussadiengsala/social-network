const CustomInput = ({
  type,
  label,
  placeholder,
  icon,
  register,
  error,
  name,
  value,
  a,
  ...additionalProps
}) => {
  return (
    <div {...additionalProps}>
      <label htmlFor={name} className="sr-only">
        {label}
      </label>
      <div className="relative">
        <input
          autoComplete="off"
          {...register(name, { ...a })}
          type={type}
          id={name}
          name={name}
          value={value}
          className={`w-full rounded-lg border-gray-100 ${
            error && "border-pink-600"
          } border-2 p-3 pe-12 text-sm shadow-sm`}
          placeholder={placeholder}
        />
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

export default CustomInput;

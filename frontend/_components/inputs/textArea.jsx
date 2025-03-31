const TextArea = ({
  name,
  additionalAttributes,
  rows,
  placeholder,
  register,
  error,
}) => {
  return (
    <div>
      <textarea
        id={name}
        {...register(name)}
        name={name}
        className={`mt-2 w-full rounded-lg border-gray-100  ${
          error && "border-pink-600"
        }  p-3 border-2 align-top shadow-sm sm:text-sm`}
        rows={rows}
        placeholder={placeholder}
        {...additionalAttributes}
      ></textarea>
      {error && (
        <div className="text-pink-500 mt-2 text-xs">*{error.message}</div>
      )}
    </div>
  );
};

export default TextArea;

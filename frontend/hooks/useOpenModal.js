import { useState } from "react";

function useOpenModal() {
  const [modals, setModals] = useState([]);

  const openModal = () => {
    setModals((prevModals) => [...prevModals, true]);
  };

  const closeModal = () => {
    setModals((prevModals) => prevModals.slice(0, -1));
  };

  const ModalComponent = ({ children }) => {
    return modals.map((isOpen, index) =>
      isOpen ? (
        <div key={index} className="modal">
          {children}
        </div>
      ) : null
    );
  };

  return {
    openModal,
    closeModal,
    ModalComponent,
  };
}

export default useOpenModal;

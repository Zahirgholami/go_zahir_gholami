const express = require('express');
const router = express.Router();
const packageController = require('../controllers/packageController');

router.get('/api/v1/packages', packageController.getAllPackages);
router.get('/api/v1/packages/:id', packageController.getPackageById);
router.post('/api/v1/packages', packageController.addPackage);
router.put('/api/v1/packages/:id', packageController.updatePackage);
router.delete('/api/v1/packages/:id', packageController.deletePackage);

module.exports = router;
package pipelinesRepo

import (
	"c2/db"
	"c2/models"

	"github.com/google/uuid"
)

func Get(pipelineID string) (pipeline models.Pipeline, err error) {
	return pipeline, db.ORM.Where("id = ?", pipelineID).First(&pipeline).Error
}

func Delete(pipelineID string) (err error) {
	return db.ORM.Where("id = ?", pipelineID).Delete(&models.Pipeline{}).Error
}

func GetMultiple() (pipelines []models.Pipeline, err error) {
	return pipelines, db.ORM.Find(&pipelines).Error
}

func Upsert(pipeline *models.Pipeline) (err error) {
	if len(pipeline.ID) == 0 {
		pipeline.ID = uuid.NewString()
	}
	return db.ORM.Save(&pipeline).Error
}

func SetSettings(pipelineID, settings string) (err error) {
	var pipe models.Pipeline
	if err := db.ORM.Table("pipelines").Where("id = ?", pipelineID).First(&pipe).Error; err != nil {
		return err
	}
	pipe.Settings = settings
	return db.ORM.Save(pipe).Error
}

func GetRun(runID string) (pipelineRun models.PipelineRun, err error) {
	return pipelineRun, db.ORM.Where("id = ?", runID).Find(&pipelineRun).Error
}

func GetRuns(pipelineID string) (pipelineRuns []models.PipelineRun, err error) {
	return pipelineRuns, db.ORM.Where("pipeline_id = ?", pipelineID).Order("executed_at DESC").Find(&pipelineRuns).Error
}

func RunUpsert(pipelineExecution *models.PipelineRun) (err error) {
	return db.ORM.Save(&pipelineExecution).Error
}
